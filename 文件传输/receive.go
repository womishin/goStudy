package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	defer listener.Close()
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//阻塞等待用户连接
	conn, err := listener.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	buf := make([]byte, 1024)
	//读取对方发送的文件名
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fileName := string(buf[:n])
	//回复ok
	conn.Write([]byte("ok"))
	//接收文件内容
	receiveFile(fileName, conn)
}

func receiveFile(fileName string, conn net.Conn) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	buf := make([]byte, 1024*4)
	//接收多少,就写入多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF || n == 0{
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err:", err)
			}
			return
		}
		//往文件写入内容
		file.Write(buf[:n])
	}
}
