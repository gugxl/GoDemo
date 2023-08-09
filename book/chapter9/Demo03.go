package main

import "fmt"

type I interface {
	f()
}
type T string

func (t T) f() {
	println("T Method")
}

type Stringer interface {
	String() string
}

func main() {
	// 类型断言
	var varI I
	varI = T("Tstring")
	if v, ok := varI.(T); ok {
		fmt.Println("varI类型的断言是：", v)
		varI.f()
	}
	//Type switch类型判断
	var value interface{}
	switch str := value.(type) {
	case string:
		fmt.Println("str类型的断言是string：", str)
	case Stringer:
		fmt.Println("str类型的断言是Stringer：", str)
	default:
		fmt.Println("value类型不在上面中：", str)
	}
	value = "类型判断检查"
	str, ok := value.(string)
	if ok {
		fmt.Printf("value类型断言结果是 %T \n", str)
	} else {
		fmt.Printf("value不是string类型", str)

	}
}
