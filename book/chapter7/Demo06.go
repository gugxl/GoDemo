package main

import (
	"fmt"
	"time"
)

func timeCost(start time.Time) {
	since := time.Since(start)
	fmt.Println(since)
}

func main() {
	defer timeCost(time.Now())
	fmt.Println("start progreame")
	time.Sleep(5 * time.Second)
	fmt.Println("end progreame")

}
