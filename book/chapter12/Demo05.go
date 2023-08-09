package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	var mutex sync.Mutex
	println("Locking G0")
	mutex.Lock()
	println("Locked G0")
	wg.Add(3)

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking(G%d)\n", i)
			mutex.Lock()
			fmt.Printf("Locked(G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unLocked(G%d)\n", i)
			wg.Done()
		}(i)
	}

	time.Sleep(time.Second * 5)
	println("ready unlock G0")
	mutex.Unlock()
	println("unlock G0")
	wg.Wait()

}
