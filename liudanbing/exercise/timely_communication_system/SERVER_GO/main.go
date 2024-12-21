package main

import (
	"SERVER_GO/server_user"
	// "timely_communication_system_server/user_mini"
)

func main() {
	// 服务器的地址
	server := server_user.NewServer("127.0.0.1", 8888)
	// 启动服务器
	server.Start()
}