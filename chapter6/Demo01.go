package main

import "fmt"

type A struct {
	Face int
}

//type Aa  A//自定义新类型Aa，没有基础类型A的方法
type Aa = A //给A取个别名

func (a A) f() {
	fmt.Println("hello", a.Face)
}

func main() {
	var s A = A{Face: 10}
	s.f()

	var sa Aa = Aa{Face: 100}
	sa.f()
}
