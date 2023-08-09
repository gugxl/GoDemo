package main

import "fmt"

var G int = 7

func N() func(int) int {
	var i int
	return func(d int) int {
		fmt.Printf("i: %d, i的地址：%p\n", i, &i)
		i += d
		return i
	}
}

func main() {
	// 全局变量G
	y := func() int {
		fmt.Printf("G:%d, G的地址是：%p\n", G, &G)
		G += 1
		return G
	}
	fmt.Println(y(), y)
	fmt.Println(y(), y)
	fmt.Println(y(), y) // y的地址

	z := func() int {
		G += 1
		return G
	}()

	// z 是直接执行的的，所以结果不会变
	fmt.Println(z, &z)
	fmt.Println(z, &z)

	// 影响外层自由变量
	var f = N()
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)

	var f1 = N()
	fmt.Println(f1(1), &f1)
}
