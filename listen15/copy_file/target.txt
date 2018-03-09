package main

import (
"os"
"github.com/astaxie/beego"
"fmt"
"io"
)

func main() {
	_, err := CopyFile("listen15/copy_file/target.txt", "listen15/copy_file/main.go")
	if err != nil {
		beego.Error(fmt.Printf("copy failed err:%v\n", err))
		return
	}
	beego.Informational("copy file OK!")
}

func CopyFile(destFile, sourceFile string) (written int64, err error){
	source, err := os.Open(sourceFile)
	if err != nil {
		beego.Error(fmt.Printf("open source file %s failed, err:%v\n", sourceFile, err))
		return
	}
	defer source.Close()
	dest, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		beego.Error(fmt.Printf("open dest file %s failed, err:%v\n", destFile, err))
		return
	}
	defer dest.Close()
	return io.Copy(dest, source)
}