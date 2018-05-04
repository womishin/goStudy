package main

import (
	"fmt"
	"os"
	"net"
	"io"
)

func main() {
	fmt.Println("请输入需要传输的文件:")
	var path string
	fmt.Scan(&path)
	//获取文件名
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	//给接收方,先发送文件名
	_, err = conn.Write([]byte(fileInfo.Name()))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	//接收对方的回复
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	if string(buf[:n]) == "ok" {
		//发送文件内容
		sendFile(path, conn)
	}
}

func sendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	//读文件内容,读多少,就给接收方发送多少
	buf := make([]byte, 1024*4)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("file.Read err:", err)
			}
			return
		}
		//发送内容
		conn.Write(buf[:n])
	}
}
