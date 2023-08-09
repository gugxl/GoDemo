package main

import "fmt"

type T struct {
	a int
}

func (t T) Mv(a int) int {
	fmt.Printf("Mv的值是: %d", a)
	return a
} // 值方法

func (tp *T) Mp(f float32) float32 {
	fmt.Printf("Mp的值是: %f", f)
	return f

} // 指针方法

func main() {
	var t T
	// 下面三种调用方式是等价的
	t.Mv(1)    // 1. 一般调用
	T.Mv(t, 1) // 2. 显示接收器调用，接收器必须是第一个
	f0 := t.Mv // 3. 选择接收器调用
	f0(2)
	T.Mv(t, 3)
	(T).Mv(t, 4)

	f := T.Mv
	f(t, 5)

}
