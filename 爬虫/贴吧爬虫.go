package main

import (
	"fmt"
	"strconv"
	"net/http"
	"log"
	"os"
)

func main() {
	var start, end int64
	fmt.Println("请输入起始页")
	fmt.Scan(&start)
	fmt.Println("请输入终止页")
	fmt.Scan(&end)
	doWork(start, end)
}

func doWork(start int64, end int64) {
	fmt.Printf("爬取%d-%d页", start, end)
	page := make(chan int64)
	for i := start; i <= end; i++ {
		go spiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		go log.Printf("%d页已爬完\n", <-page)
	}
}

func spiderPage(i int64, page chan<- int64) {
	url := "http://tieba.baidu.com/f?kw=%E5%9B%BE%E6%8B%89%E4%B8%81&ie=utf-8&pn=" + strconv.FormatInt((i-1)*50, 10)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	writeFile(resp, i)
	page <- i
}

func writeFile(resp *http.Response, i int64) {
	fileName := strconv.FormatInt(i, 10) + ".html"
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	buf := make([]byte, 1024*4)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if err != nil || n == 0 {
			break
		}
		result += string(buf[:n])
	}
	file.WriteString(result)
}
