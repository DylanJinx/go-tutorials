package main

import (
	"fmt"
	"net"
	"sync"
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
		conn, err := listener.Accept()  // 当accept成功，代表有一个客户端连接进来
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
	user := NewUser(conn)

	  // 用户上线了，将用户加入到OnlineMap中
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()


	  // 广播当前用户上线消息
	s.BroadCast(user, "已上线")

	  // 让当前的handler一直阻塞
	select {}
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