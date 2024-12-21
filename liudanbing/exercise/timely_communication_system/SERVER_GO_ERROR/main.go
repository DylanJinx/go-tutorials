package main

import (
	"SERVER_GO_ERROR/server_mini"
	// "timely_communication_system_server/user_mini"
)

func main() {
	// 服务器的地址
	server := server_mini.NewServer("127.0.0.1", 8888)
	// 启动服务器
	server.Start()
}