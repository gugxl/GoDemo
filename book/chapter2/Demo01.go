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
}
