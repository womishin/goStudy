package iniconfig

import (
	"io/ioutil"
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T){
	bytes, err := ioutil.ReadFile("c:/config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))

	err = ioutil.WriteFile("d:/config.ini", bytes, 0755)
	if err != nil {
		fmt.Println(err)
	}
}
