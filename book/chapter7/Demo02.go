package main

import "fmt"

func div(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获异常 %s\n", r)
		}
	}()

	if b < 0 {
		panic("除数需要大于0")
	}

	fmt.Println("余数为：", a/b)
}

func main() {
	div(5, 1)

	defer func() { // 拦截本方法里面的异常
		fmt.Println("done")           // 在这个前面出现异常，不会打印这行，后面出现会打印
		if r := recover(); r != nil { // 接受本方法内异常并处理
			fmt.Printf("run time panic %v \n", r)
		}

	}()

	fmt.Println("start")
	panic("Error") // 抛出异常

}
