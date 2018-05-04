package 乞丐版并发服务器

import (
	"os"
	"fmt"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("useage :xxx file")
		return
	}
	fileName := list[1]
	fileInfo, err := os.Stat(fileName)
	if err != nil{
		fmt.Println("err=", err)
		return
	}
	fmt.Println("name=", fileInfo.Name())
	fmt.Println("size=", fileInfo.Size())
}
