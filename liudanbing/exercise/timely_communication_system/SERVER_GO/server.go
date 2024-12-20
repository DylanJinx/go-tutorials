package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip string
	Port int

	                            // 在线用户的列表
	OnlineMap map[string]*User  // key: Name, value: *User
	mapLock   sync.RWMutex      // 关于OnlineMap的读写锁

	  // 消息广播的channel
	Message chan string
}

  // 创建一个server的接口
func NewServer(ip string, port int) *Server {
	  // 此时的server是一个指针
	server := &Server{
		Ip       : ip,
		Port     : port,
		OnlineMap: make(map[string]*User),
		Message  : make(chan string),
	}

	return server
}

  // 启动服务器的接口，是Server的方法, S大写表示public
func (s *Server) Start() {
	  // socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if       err  != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	  // close listen socket
	defer listener.Close()

	  // 启动监听Message的goroutine
	go s.ListenMessage()

	for {
		                                // accept
		conn, err := listener.Accept()  // 当accept成功，代表有一个客户端连接进来，conn是和客户端通信的接口
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		  // do handler
		go s.Handler(conn)
	}

}

func (s *Server) Handler(conn net.Conn) {
	  // 当前连接的业务
	  // fmt.Println("连接建立成功")

	  // 创建一个用户
	user := NewUser(conn, s)

	/*v3 -> v4
	  // 用户上线了，将用户加入到OnlineMap中
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()
	  // 广播当前用户上线消息
	s.BroadCast(user, "已上线")
	*/
	user.Online() // v4

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf) // n代表读取到的字节数
			if n == 0 { // n == 0代表客户端断开
				/*v3 -> v4
				s.BroadCast(user, "下线")  // 广播用户下线消息
				*/ 
				user.Offline()  // v4
				return
			}

			if err != nil && err != io.EOF {  // io.EOF代表读到文件末尾
				fmt.Println("conn.Read err:", err)
				return
			}

			// 提取用户的消息(去除\n)
			msg := string(buf[:n-1])  // 将读取到的字节转换成字符串

			/*v3 -> v4
			// 将得到的消息进行广播
			s.BroadCast(user, msg)
			*/
			// 用户针对msg进行处理
			user.DoMessage(msg)  // v4

			// 用户的任意消息，代表当前用户是活跃的
			isLive <- true
		}
	}()

	for {
		select {
		case <- isLive:
			// 当前用户是活跃的，应该重置定时器
			// 不做任何事情，为了激活select，更新下面的定时器

		case <- time.After(time.Second * 5):
			// 已经超时
			// 将当前的user强制关闭

			user.SendMessage("你被踢了")
			// 销毁用户的goroutine
			close(user.C)
			// 关闭连接
			conn.Close()
			// 退出当前的handler
			return // runtime.Goexit()
		}
	}
}

  // 广播消息的方法(arg1: 由哪个用户发起的, arg2: 消息内容)
func (s *Server) BroadCast(user *User, msg string) {
	sandMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	s.Message <- sandMsg  // 将消息发送到Message channel中
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部在线的用户
func (s *Server) ListenMessage() {
	for {
		msg := <- s.Message

		  // 将msg发送给全部在线用户
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg  // 将消息发送到用户的channel中
		}

		s.mapLock.Unlock()
	}
}