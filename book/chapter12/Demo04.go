package main

import (
	"sync"
	"time"
)

var mutex sync.Mutex

func LockA() {
	mutex.Lock()
	println("Lock in A")
	LockB()
	time.Sleep(5)
	println("Wake up in A")
	mutex.Unlock()
	println("UnLock in A")

}

func LockB() {
	println("B")
	mutex.Lock()
	println("Lock in B")
	mutex.Unlock()
	println("UnLock in B")
}

func main() {
	LockA()
	time.Sleep(10)
}
