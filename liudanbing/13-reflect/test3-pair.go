package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	// tty: pair<type: *osFile, value: "/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0) // os.OpenFile() 函数返回一个文件对象和一个错误对象，打开 /dev/tty 设备文件，以读写模式打开
	tty.Write([]byte("bye\n")) // 将字符串 "bye" 写入到 /dev/tty 设备文件中

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// r: pair<type: , value: >
	var r io.Reader // io.Reader 是一个接口类型
	// r pair<type: *osFile, value: "/dev/tty"文件描述符>
	r = tty

	// w: pair<type: , value: >
	var w io.Writer // io.Writer 是一个接口类型
	// w pair<type: *osFile, value: "/dev/tty"文件描述符>
	w = r.(io.Writer) // 将 r 的值断言为 io.Writer 类型，然后赋值给 w

	w.Write([]byte("hello")) // 调用 w 的 Write() 方法，将字符串 "hello" 写入到 /dev/tty 设备文件中
}