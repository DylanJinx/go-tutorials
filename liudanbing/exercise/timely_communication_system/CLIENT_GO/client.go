package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

var serverIp  string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

type Client struct {
	ServerIp   string
	ServerPort int
	Name	   string
	conn       net.Conn
	flag       int            // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client {
		ServerIp  : serverIp,
		ServerPort: serverPort,
		flag      : 999,
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

func (c *Client) menu() bool {
	fmt.Println(">>>>>> 1. 公聊模式")
	fmt.Println(">>>>>> 2. 私聊模式")
	fmt.Println(">>>>>> 3. 更新用户名")
	fmt.Println(">>>>>> 4. 查询在线用户")
	fmt.Println(">>>>>> 0. 退出")

	/*
	var flag int

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>>>>> 请输入合法范围内的数字")
		return false
	}
	*/

	// 优化：可以处理字符串输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入读取内容
	input, err := reader.ReadString('\n') // 读取直到遇到\n
	if err != nil {
		fmt.Println("reader.ReadString err:", err)
		return false
	}

	input = strings.TrimSpace(input) // 去掉input两端的空格
	flag, err := strconv.Atoi(input) // 将input转换成int类型
	if err != nil {
		fmt.Println(">>>>>>请输入合法范围内的数字")
		return false
	}

	if flag >= 0 && flag <= 4 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>请输入合法范围内的数字")
		return false
	}
}

func (c *Client) UpdateName() bool {
	fmt.Println(">>>>>> 请输入用户名:")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([]byte(sendMsg)) // 将sendMsg发送给服务器
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

// 公聊模式
func (c *Client) PublicChat() {
	fmt.Println(">>>>>> 请输入聊天内容，exit退出")
	var chatMsg string
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := c.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>> 请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (c *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}

// 私聊模式
func (c *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	c.SelectUsers()

	fmt.Println(">>>>>> 请输入聊天对象用户名，exit退出:")
	fmt.Scanln(&remoteName)

	// 选择聊天对象
	for remoteName != "exit" {
		fmt.Println(">>>>>> 请输入消息内容，exit退出:")
		fmt.Scanln(&chatMsg)

		// 与这个聊天对象进行聊天
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := c.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write err:", err)
					break
				}
			}

			// 多次输入消息
			// chatMsg = ""
			fmt.Println(">>>>>> 请输入消息内容，exit退出:")
			fmt.Scanln(&chatMsg)
		}

		// 重新选择聊天对象
		//remoteName = ""
		c.SelectUsers()
		fmt.Println(">>>>>> 请输入聊天对象用户名，exit退出:")
		fmt.Scanln(&remoteName)

	}
}

func (c *Client) Run() {
	for c.flag != 0 { // 不断得循环，判断flag是否为0
		for !c.menu() {} // 如果menu()返回false，那么就一直循环，直到menu()返回true

		// 根据不同的模式处理不同的业务
		switch c.flag { // switch会自动break，如果需要继续执行下一个case，需要使用fallthrough，但是fallthrough仅仅会转到下一个case
		case 1:
			// 公聊模式
			fmt.Println(">>>>>> 公聊模式")
			c.PublicChat()

		case 2:
			// 私聊模式
			fmt.Println(">>>>>> 私聊模式")
			c.PrivateChat()

		case 3:
			// 更新用户名
			fmt.Println(">>>>>> 更新用户名")
			c.UpdateName()
		case 4:
			// 查询在线用户
			fmt.Println(">>>>>> 查询在线用户")
			c.SelectUsers()
		}
	}
}

// 这段逻辑不能写到Run()中，如果写到Run()中，那么Run()就会阻塞在这里，无法继续执行
func (c *Client) DealResponse() {
	// 一旦client.conn有数据，就直接拷贝到os.Stdout标准输出上，永久阻塞监听
	_, err := io.Copy(os.Stdout, c.conn)
	// 等价于：
	// for {
	// 	buf := make([]byte, 4096)
	// 	n, err := c.conn.Read(buf)
	// 	fmt.Println(buf[:n])

	if err != nil {
		fmt.Println("\n>>>>>> 与服务器的连接已断开，客户端即将退出...")
	}
	// 无论是正常EOF还是错误导致的退出，都结束客户端
	os.Exit(0)
}


func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>> 连接服务器失败")
		return
	}

	// 单独开启一个goroutine处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>> 连接服务器成功")

	client.Run()
}