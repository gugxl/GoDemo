package main

import "fmt"

var (
	Ga int = 99
)

const (
	v int = 199
)

func GetGa() func() int {
	if Ga := 55; Ga < 60 {
		fmt.Println("if中Ga:", Ga)
	}

	for Ga := 2; ; {
		fmt.Println("for中Ga:", Ga)
		break
	}
	fmt.Println("GetGa函数中Ga:", Ga)

	return func() int {
		Ga += 1
		return Ga
	}
}

func main() {
	Ga := "string"
	fmt.Println("main函数中Ga：", Ga)

	b := GetGa()
	fmt.Println("main函数中：", b(), b(), b(), b(), b(), b())
	v := 1
	{
		v := 2
		println(v)
		{
			v := 3
			println(v)
		}
	}
	println(v)

	if a := 1; false {

	} else if b := 2; false {

	} else if c := 3; false {

	} else {
		println(a, b, c)
	}

}
