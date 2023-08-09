package main

import "fmt"

func main() {
	var i int = 1

	// defer 声明时会实时解析[直接计算]，执行顺序是先进后出[先执行的后返回]
	defer fmt.Println("result1 =>", func() int { return i * 2 }())
	i++

	defer func() {
		fmt.Println("result2 =>", i*2)
	}()
}
