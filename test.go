package main

import (
	"net/http"
	"log"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024*4)
	var tmp string
	for{
		n, err := resp.Body.Read(buf)
		if err != nil || n == 0{
			log.Println(err)
			break
		}
		tmp += string(buf[:n])
	}
	log.Println(tmp)
}
