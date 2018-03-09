package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogFuncCall(true)
}

func main() {
	fmt.Println(funcA())
}

func funcA() int {
	x := 5
	defer func() {
		x += 1
		beego.Informational(x)
	}()
	beego.Informational(x)
	return x
}
