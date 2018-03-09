package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	str := "Hello World!"
	err := ioutil.WriteFile("listen15/ioutil_write/test.dat", []byte(str), 0755)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
}
