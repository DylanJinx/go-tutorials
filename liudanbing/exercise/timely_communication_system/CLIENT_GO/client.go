package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name	   string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client {
		ServerIp  : serverIp,
		ServerPort: serverPort,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", client.ServerIp, client.ServerPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}

	client.conn = conn

	// 返回对象	
	return client
}

var serverIp  string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	// 命令行解析
	flag.Parse()
	
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>> 连接服务器失败")
		return
	}

	fmt.Println(">>>>>> 连接服务器成功")

	// select {} // 阻塞

	for {
		time.Sleep(1 * time.Second)
	}
}