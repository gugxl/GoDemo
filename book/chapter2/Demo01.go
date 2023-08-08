package main

import "fmt"

const s = "Go 语言"

func main() {
	for i, r := range s {
		fmt.Printf("%#U : %d\n", r, i)
	}

	// 虚数， 啊啊啊啊，什么鬼
	var c complex128 = 1.0 + 10i
	fmt.Printf("%v", c)

	var cc = complex(5.67, 99.09)
	fmt.Printf("re: ", real(cc), "im:", imag(cc))

	//var (
	//	a   int
	//	b   bool
	//	str string
	//	浮点  float32
	//)
	//println(a)
	//println(b)
	//println(str)
	//println(浮点)

	a, b, x := 5, true, "aasa"
	println(a)
	println(b)
	println(x)

}
