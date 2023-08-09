package main

import "fmt"

func fun1() (i int) {
	defer func() {
		i = i + 10
	}()
	// 会被defer 返回 可以写 return i 或者 return
	return 0
}

func main() {
	var i int = 1

	// defer 声明时会实时解析，执行顺序是先进后
	// 先执行return ，暂存return值，按照定义把 defer 压入栈 逆序执行
	defer fmt.Println("result1 =>", func() int { return i * 2 }())

	i++

	defer func() {
		fmt.Println("result2 =>", i*2)
	}()
	i++

	//打印 10
	fmt.Println("result3 =>", fun1())

}
