package main

import "fmt"

func Add(a, b int) {
	fmt.Printf("%d + %d = %d", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
func main() {
	callback(1, Add)

	// 匿名函数
	f := func(x, y int) int {
		return x + y
	}
	println(f(3, 5))

	func() {
		sum := 0
		for i := 0; i <= 1e6; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()

}
