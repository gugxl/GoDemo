package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	go hello()
	select {}
}

func hello() {
	for true {
		go func() {
			fmt.Println("hello world")
		}()
		time.Sleep(time.Millisecond)
	}
}

// 安装graphviz bin 加入path后
// go tool pprof http://localhost:8080/debug/pprof/heap
// 输入 png 会下载图片
