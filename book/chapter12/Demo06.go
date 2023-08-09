package main

import (
	"fmt"
	"sync"
	"time"
)

type Book struct {
	BookName string
	L        *sync.Mutex
}

func (bk *Book) setName(wg *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name", name)
		bk.L.Unlock()
		wg.Done()
	}()

	bk.L.Lock()
	fmt.Println("lock set name", name)

	time.Sleep(time.Second)
	bk.BookName = name
}

func main() {
	bk := Book{}
	bk.L = new(sync.Mutex)
	wg := &sync.WaitGroup{}
	books := []string{"三国演义", "水浒传", "西游记"}
	for _, book := range books {
		wg.Add(1)
		go bk.setName(wg, book)
	}
	wg.Wait()
}
