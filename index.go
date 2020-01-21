package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//创建文件
	now := time.Now()
	parallelism := flag.Bool("p", false, "Whether to use parallelism")
	filename := flag.String("name", "default", "filename")
	Shense := flag.Int("s", 6, "use sb shense")
	flag.Parse()
	//PathExists("www")
	if *parallelism {
		saveMavP(loadtxt(*filename), strconv.Itoa(*Shense), *filename)
	} else {
		saveMav(loadtxt(*filename), strconv.Itoa(*Shense), *filename)
	}
	finish2 := time.Since(now)
	fmt.Println("总花费时间：", finish2)
}

// PathExists ...
func PathExists(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir("./"+path, os.ModePerm)
	}
}

func saveMavP(send []string, ss string, filename string) {
	//创建文件
	PathExists(filename)
	page := make(chan string)
	//分片逐步写入
	for i, char := range send {
		go SpiderPage(char, page, i, filename, ss)
		time.Sleep(time.Second / 10)
	}
	for i := 0; i < len(send); i++ {
		fmt.Println("下载完成", <-page)
	}
}

// SpiderPage ...
func SpiderPage(char string, page chan string, i int, filename string, ss string) {
	path := "./" + filename + "/" + strconv.Itoa(i+1) + ".mp3"
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	r, _ := http.Get("https://ai.qq.com/cgi-bin/wxappdemo_ttsecho?text=" + char + "&speaker=" + ss + "&speed=100&volume=0&format=3&aht=0&apc=58&download=1")
	defer func() { _ = r.Body.Close() }()
	b, _ := ioutil.ReadAll(r.Body)
	f.Write(b)
	page <- path
}

func saveMav(send []string, ss string, filename string) {
	//创建文件
	path := "./" + filename + ".mp3"
	f, err := os.Create(path)
	var chanStream bytes.Buffer
	if err != nil {
		return
	}
	defer f.Close()
	//分片逐步写入
	for _, char := range send {
		r, _ := http.Get("https://ai.qq.com/cgi-bin/wxappdemo_ttsecho?text=" + char + "&speaker=" + ss + "&speed=100&volume=0&format=3&aht=0&apc=58&download=1")
		defer func() { _ = r.Body.Close() }()
		b, _ := ioutil.ReadAll(r.Body)
		chanStream.Write(b)
	}
	f.Write(chanStream.Bytes())
}

func loadtxt(filename string) []string {
	txt, _ := os.OpenFile("./"+filename+".txt", os.O_RDONLY, 0600)
	defer txt.Close()
	contentByte, _ := ioutil.ReadAll(txt)
	return strings.Split(strings.Replace(string(contentByte), "\r\n", "\n", -1), "\n")
}
