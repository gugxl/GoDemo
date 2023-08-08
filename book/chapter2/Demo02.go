package main

import "fmt"

func main() {
	x := 1
	fmt.Print(x)
	{
		fmt.Print(x)
		// 使用 := 相当于临时变量, 不会影响外面的值，使用=会赋值给外面
		x := "asd"
		fmt.Print(x)
	}
	fmt.Println(x)

	var a, b = 1, 2
	println(a, b)
	// 交换变量
	a, b = b, a
	println(a, b)

	// _ 只写变量，这个变量的值没法获取，相当于丢掉
	var _, c = 5, 3
	print(c)
}
