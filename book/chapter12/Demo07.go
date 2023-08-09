package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	var rwMutex sync.RWMutex
	Data := 0
	for i := 0; i < 10; i++ {
		go func(t int) {
			rwMutex.RLocker()
			defer rwMutex.RLocker()
			fmt.Println("读数据：%v %d \n", Data, i)
			wg.Done()
			time.Sleep(time.Second)
			// 第一次运行后读解锁
			// 循环到第二个的时候，读锁定后，这个go routine就没有阻塞，同时读成功

		}(i)

		go func(t int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			Data += t
			fmt.Println("写数据：%v %d \n", Data, i)
			wg.Done()
			// 对读写锁进行读锁定或者写锁定都会阻塞。写锁定下是需要写解锁之后才能写的
			time.Sleep(5 * time.Second)
		}(i)
	}
	wg.Wait()
}
