package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	if _, ok := m["one"]; !ok {
		fmt.Println("no exist")
	} else {
		fmt.Println(" exist")
	}

	// 不会改变原来数组内容
	data := []int{1, 2, 3}
	for i := range data {
		i *= 10
	}
	fmt.Println("data:", data)

	// 下标访问修改原值
	for i, _ := range data {
		data[i] *= 10
	}
	fmt.Println("data:", data)

}
