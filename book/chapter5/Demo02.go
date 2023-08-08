package main

import "fmt"

func main() {
	sli := make([]int, 5, 10)
	fmt.Println("sli切片的长度和容量：%d,%d", len(sli), cap(sli))
	fmt.Println(sli)
	newSli := sli[:cap(sli)]
	fmt.Println(newSli)

	var x = []int{1, 3, 5, 8, 11}
	fmt.Println("x切片的长度和容量：%d,%d", len(x), cap(x))

	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:4] // a[low:high:max]  max - low 是容量，high-low是长度
	fmt.Println("t切片的长度和容量：%d,%d", len(t), cap(t))
	fmt.Println(t)
	//fmt.Println(t[2])     // panic 索引不能超过切片长度

}
