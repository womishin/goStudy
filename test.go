package main

import (
	"runtime"
	"fmt"
)

func main() {
	n := runtime.GOMAXPROCS(100)
	fmt.Println("核心数:", n)

	for{
		go fmt.Print(0)
		fmt.Print(1)
	}
}
