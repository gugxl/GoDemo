package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5) // 5是缓存区大小，如果不带的话相当于单条。
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ready", i)
		c <- i
		fmt.Println("send", i)
	}
}

func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received", i)
	}
}
