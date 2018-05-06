package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用于发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

var (
	onLineMap = make(map[string]Client)
	message   = make(chan string)
	isQuit    = make(chan bool) //对方是否退出
	hasDate   = make(chan bool) //对方是否有数据发送
)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	defer listener.Close()
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//新开一个协程,转发消息,只要有消息来了,遍历map,给map中的每个成员都发送消息
	go manager()
	//主协程,循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		//处理用户连接
		go handleConn(conn)
	}
}

func manager() {
	for {
		//没有消息前,这里会阻塞
		msg := <-message
		for _, cli := range onLineMap {
			cli.C <- msg
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()
	cli := Client{make(chan string), cliAddr, cliAddr}
	onLineMap[cliAddr] = cli
	//新开一个协程,给当前客户端发送信息
	go writeMsgToClient(cli, conn)
	//广播某个在线
	message <- makeMsg(cli, "login")
	//提示,我是谁
	cli.C <- makeMsg(cli, "I am here")
	//新建一个协程,接收用户发送过来的信息
	go func() {
		for {
			buf := make([]byte, 2048)
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("conn.Read err:", err)
				return
			}
			msg := string(buf[:n-1])
			if len(msg) == 3 && msg == "who" {
				//遍历map,给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onLineMap {
					msg := tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				cli.Name = strings.Split(msg, "|")[1]
				onLineMap[cliAddr] = cli
				conn.Write([]byte("rename ok!\n"))
			} else {
				message <- makeMsg(cli, msg)
			}
			hasDate <- true
		}
	}()
	for {
		//通过select检测channel的流动
		select {
		case <-isQuit:
			//移除当前用户
			delete(onLineMap, cliAddr)
			//广播谁下线了
			message <- makeMsg(cli, "login out")
		case <-hasDate:
		case <-time.After(60 * time.Second):
			delete(onLineMap, cliAddr)
			message <- makeMsg(cli, "time out leave out")
			return
		}
	}
}

func makeMsg(cli Client, msg string) string {
	return "[" + cli.Addr + "]" + cli.Name + ":" + msg
}

func writeMsgToClient(cli Client, conn net.Conn) {
	conn.Close()
	for msg := range cli.C {
		//给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}
