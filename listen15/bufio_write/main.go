package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	file, err := os.OpenFile("listen15/bufio_write/test.dat", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello World!\n")
	}
	writer.Flush()
}
