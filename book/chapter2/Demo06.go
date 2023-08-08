package main

import "strconv"

func main() {
	//fallthrough 这个不会判断下一个分支的条件会直接执行
	switch {
	case false:
		println("false")
		fallthrough
	case true:
		println("true")
		fallthrough
	case false:
		println("false fallthrough  ")
		fallthrough
	case true:
		println("true fallthrough")
		fallthrough
	default:
		println("default")
	}

	switch a := 1; {
	case a == 1:
		println("a ==1")
		fallthrough
	case a == 2:
		println("a == 2")
	case a == 3:
		println("a == 3")
		fallthrough
	case a == 4:
		println("a == 4")
		fallthrough
	default:
		println("default a==" + strconv.Itoa(a))
	}

}
