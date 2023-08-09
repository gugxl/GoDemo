package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) { // go 关键字直接创建了一个协程
			println(i)
		}(i)
	}

	time.Sleep(1e9)
}
