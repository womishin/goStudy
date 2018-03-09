package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.OpenFile("listen15/write/test.dat", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	file.WriteString("Hello World!")
}
