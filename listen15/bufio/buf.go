package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func main() {
	//只读的方式打开
	file, err := os.Open("listen15/bufio/buf.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Println(line)
	}
}

