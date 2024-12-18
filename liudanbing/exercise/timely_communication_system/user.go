package main

import "net"

type User struct {
	Name string 
	Addr string 
	C    chan string  // 和用户绑定的channel
	conn net.Conn     // 是用户唯一可以和对端客户端通信的接口
}

  // 创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()  // 获取远程客户端的地址
	user     := &User {
		Name: userAddr,
		Addr: userAddr,
		C   : make(chan string),
		conn: conn,
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