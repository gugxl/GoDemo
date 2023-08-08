package main

import (
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	println(p.name)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for j := 0; j < 5; j++ {
		for i := 0; i < 6; i++ {
			println(i)
		}
	}

	data := []field{{"one"}, {"two"}, {"three"}}
	// v是指针 下面这种情况会导致指针修改值，输出 three three three
	for _, v := range data {
		go v.print()
	}
	time.Sleep(3 * time.Second)

	// v 是普通值的副本，不会修改原来对象
	data2 := []string{"one", "two", "three"}
	for _, v := range data2 {
		go func(in string) {
			println(in)
		}(v)
	}
	time.Sleep(time.Second * 3)

}
