package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	// 此时的server是一个指针
	server := &Server{
		Ip: ip,
		Port: port,
	}

	return server
}

// 启动服务器的接口，是Server的方法, S大写表示public
func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
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
	fmt.Println("连接建立成功")
}