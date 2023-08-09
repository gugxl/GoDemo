package main

import (
	"fmt"
	"time"
)

func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return 0
}

type funcType func(time time.Time)

func main() {

	f := func() int { return 7 }
	println(f())

	f1 := func(t time.Time) time.Time { return t } // 方式1直接赋值给变量
	fmt.Println(f1(time.Now()))

	var timer funcType = CurrentTime // 方式2 定义函数类型funcType变量time
	timer(time.Now())

	funcType(CurrentTime)(time.Now()) // 方式3 先把CurrentTime转换funcType，然后传入参数进行调用

	// 常规调用参数传多个
	list(1, 2, 3, 4)

	num := []int{1, 3, 5, 6, 7}
	list(num...)
}

func list(nums ...int) {
	fmt.Println(nums)
}

func CurrentTime(start time.Time) {
	fmt.Println(start)
}
