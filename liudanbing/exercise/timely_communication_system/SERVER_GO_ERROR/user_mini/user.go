package user_mini

import (
	"SERVER_GO_ERROR/server_mini"
	"net"
	"strconv"
	"strings"
)

type User struct {
	Name string 
	Addr string 
	C    chan string  // 和用户绑定的channel
	conn net.Conn     // 是用户唯一可以和对端客户端通信的接口

	server *server_mini.Server // 当前用户所在的server
}

  // 创建一个用户的API
func NewUser(conn net.Conn, server *server_mini.Server) *User {
	userAddr := conn.RemoteAddr().String()  // 获取远程客户端的地址
	user     := &User {
		Name: userAddr,
		Addr: userAddr,
		C   : make(chan string),
		conn: conn,

		server: server,
	}

	  // 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

  // 每个user都应该启动一个goroutine来处理server的消息，即监控channel，如果有消息就发送给客户端
func (u *User) ListenMessage() {
	for {
		msg := <- u.C

		u.conn.Write([]byte(msg + "\n"))  // 这行是意思是将msg + 转义字符\n 转换成byte类型，然后写入到u.conn中，即发送给客户端
	}
}

// 用户上线的业务
func (u *User) Online() {
	// 用户上线了，将用户加入到OnlineMap中
	u.server.MapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.MapLock.Unlock()

	// 广播当前用户上线消息
	u.server.BroadCast(u, "已上线")
}

// 用户下线的业务
func (u *User) Offline() {
	// 用户下线，将用户从OnlineMap中删除
	u.server.MapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.MapLock.Unlock()

	// 广播当前用户下线
	u.server.BroadCast(u, "已下线")
}

// 用户处理消息的业务
func (u *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户有哪些

		u.server.MapLock.Lock()
		i := 1

		for _, user := range u.server.OnlineMap {
			onlineMsg := strconv.Itoa(i) + ":" + "[" + user.Addr + "]" + user.Name + ":" + "在线\n"
			u.SendMessage(onlineMsg) // 或者 u.C <- onlineMsg
			i++
		}

		u.server.MapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // msg[:7]是取msg的前7个字符
		// 消息格式：rename|张三
		newName := strings.Split(msg, "|")[1]  // 通过|分割msg，取第二个元素；或者使用msg[7:]来取msg的第8个字符到最后一个字符
		// 判断newName是否存在
		_, ok := u.server.OnlineMap[newName]
		if ok {
			u.SendMessage("当前用户名被使用\n") // 或者 u.C <- "当前用户名被使用\n"
		} else {
			u.server.MapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.Name = newName
			u.server.OnlineMap[newName] = u
			u.server.MapLock.Unlock()

			u.SendMessage("您已经更新用户名:" + u.Name + "\n") // 或者 u.C <- "您已经更新用户名:" + u.Name + "\n"
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式：to|张三|消息内容
		
		// 1. 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			u.SendMessage("消息格式不正确，请使用\"to|张三|消息内容\"\n")
			return
		}
		// 2. 根据用户名得到对方User对象
		remoteUser, ok := u.server.OnlineMap[remoteName]
		if !ok {
			u.SendMessage("该用户名不存在\n")
			return
		} 

		// 3. 获取消息内容，通过对方的User对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.SendMessage("无消息内容，请重发\n")
			return
		}

		remoteUser.SendMessage(u.Name + "对您说：" + content + "\n")
	} else {
		// 将用户发送的消息进行广播
		u.server.BroadCast(u, msg)

	}

}

// 给当前用户的客户端发送消息
func (u *User) SendMessage(msg string) {
	u.conn.Write([]byte(msg))
}