package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr)
	}(&x)
	fmt.Println(x)

	//var x = nil 编译失败

	m := make(map[string]int)
	m["one"] = 1
	print(m["one"])

	// 常量
	const b string = "abc"
	const c = "xxxx"

}
