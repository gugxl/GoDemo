package main

import "fmt"

func main() {
	fmt.Println("fun1 return: ", func1())   //102
	fmt.Println("fun2 return: ", func2())   // 返回暂存值
	fmt.Println("fun3 return: ", func3())   // 5
	fmt.Println("fun33 return: ", func33()) // 5
	fmt.Println("fun4 return: ", func4())
}
func func1() (i int) {
	defer func() {
		i++
		fmt.Println("fun1 defer2: ", i) // 101
	}()

	defer func() {
		i++
		fmt.Println("fun1 defer1: ", i) // 102
	}()
	// 这个值会赋值给i
	return 100
}

func func2() int {
	var i int
	defer func() {
		i++
		fmt.Println("fun2 defer2: ", i) // 2
	}()

	defer func() {
		i++
		fmt.Println("fun2 defer1: ", i) // 1
	}()
	// 这个值会暂存，跟i无关
	return i
}

func func3() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println("fun3 defer: ", t) // 10
	}()
	return t
}

// 等价func3
func func33() (r int) {
	r = 5
	t := 5
	func() {
		t = t + 5
		fmt.Println("fun33:", t)
	}()
	return
}

func func4() int {
	i := 8
	defer func(i int) {
		fmt.Println("fun4 defer: ", i) // 8
	}(i)
	i = 19
	return i // 暂存19
}
