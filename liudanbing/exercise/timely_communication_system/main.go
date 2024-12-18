package main

func main() {
	// 服务器的地址
	server := NewServer("127.0.0.1", 8888)
	// 启动服务器
	server.Start()
}