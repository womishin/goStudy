package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"io"
	"regexp"
	"strings"
	"os"
)

func main() {
	var start, end int64
	fmt.Println("请输入起始页")
	fmt.Scan(&start)
	fmt.Println("请输入终止页")
	fmt.Scan(&end)
	doWork1(start, end)
}

func doWork1(start int64, end int64) {
	fmt.Printf("爬取%d-%d页的数据\n", start, end)
	page := make(chan int64)
	for i := start; i <= end; i++ {
		go spiderPage1(i, page)
	}
	for i := start; i <= end; i++ {
		go log.Printf("第%d页已爬完", <-page)
	}
}

func spiderPage1(i int64, page chan<- int64) {
	url := fmt.Sprintf("https://www.pengfu.com/xiaohua_%s.html", strconv.FormatInt(i, 10))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	makeData(resp, i)
	page <- i
}

func makeData(resp *http.Response, i int64) {
	defer resp.Body.Close()
	buf := make([]byte, 1024*4)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF || n == 0 {
			break
		}
		result += string(buf[:n])
	}
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		log.Fatalln("regexp.MustCompile err")
	}
	joyUrls := re.FindAllStringSubmatch(result, -1)
	titleSlice := make([]string, 0)
	contentSlice := make([]string, 0)
	for _, url := range joyUrls {
		title, content := spiderOneJoy(url[1])
		titleSlice = append(titleSlice, title)
		contentSlice = append(contentSlice, content)
		writeFile1(i, titleSlice, contentSlice)
	}

}

func writeFile1(i int64, titleSlice []string, contentSlice []string) {
	file, err := os.Create(strconv.FormatInt(i, 10) + ".txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < len(titleSlice); i++ {
		file.WriteString(titleSlice[i])
		file.WriteString(contentSlice[i] + "\n")
		file.WriteString("-------------------------------------------------\n\n")
	}
}

func spiderOneJoy(url string) (title, content string) {
	resp, err := http.Get(url)
	buf := make([]byte, 1024*4)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF || n == 0 {
			break
		}
		result += string(buf[:n])
	}
	if err != nil {
		log.Fatalln(err)
	}
	re := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re == nil {
		log.Fatalln("regexp.MustCompile err")
	}
	tmpTitle := re.FindAllStringSubmatch(result, 1)
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		break
	}
	re = regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re == nil {
		log.Fatalln("regexp.MustCompile err")
	}
	tmpContent := re.FindAllStringSubmatch(result, 1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		break
	}
	return
}
