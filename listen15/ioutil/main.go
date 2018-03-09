package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	content, err := ioutil.ReadFile("listen15/ioutil/main.go")
	if err != nil{
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}