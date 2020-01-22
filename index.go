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
	"sync"
	"time"
)

//SafeMap ...
type SafeMap struct {
	sync.RWMutex
	Map map[int][]byte
}

func main() {
	//创建文件
	now := time.Now()
	filename := flag.String("name", "default", "filename")
	Shense := flag.Int("s", 6, "use sb shense")
	flag.Parse()
	//PathExists("www")
	saveMavP(loadtxt(*filename), strconv.Itoa(*Shense), *filename)
	finish2 := time.Since(now)
	fmt.Println("总花费时间：", finish2)
}

func newSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int][]byte)
	return sm

}

func (sm *SafeMap) writeMap(key int, value []byte) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}

func saveMavP(send []string, ss string, filename string) {
	//创建文件
	path := "./" + filename + ".mp3"
	f, err := os.Create(path)
	if err != nil {
		return
	}
	Zongfile := newSafeMap()
	var chanStream bytes.Buffer
	page := make(chan int)
	//分片逐步写入
	for i, char := range send {
		go SpiderPage(char, page, i, ss, Zongfile)
	}
	for i := 0; i < len(send); i++ {
		fmt.Println("下载完成", <-page)
	}
	for i := 0; i < len(send); i++ {
		chanStream.Write(Zongfile.Map[i])
	}
	f.Write(chanStream.Bytes())
}

// SpiderPage ...
func SpiderPage(char string, page chan int, i int, ss string, zongfile *SafeMap) {
	r, err := http.Get("https://ai.qq.com/cgi-bin/wxappdemo_ttsecho?text=" + char + "&speaker=" + ss + "&speed=100&volume=0&format=3&aht=0&apc=58&download=1")
	if err != nil {
		fmt.Println("下载失败", i, "正在重试")
		go SpiderPage(char, page, i, ss, zongfile)
	} else {
		defer r.Body.Close()
		b, _ := ioutil.ReadAll(r.Body)
		zongfile.writeMap(i, b)
		page <- i
	}
}

func loadtxt(filename string) []string {
	txt, _ := os.OpenFile("./"+filename+".txt", os.O_RDONLY, 0600)
	defer txt.Close()
	contentByte, _ := ioutil.ReadAll(txt)
	return strings.Split(strings.Replace(string(contentByte), "\r\n", "\n", -1), "\n")
}
