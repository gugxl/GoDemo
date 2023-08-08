package main

import "fmt"

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 10000 10000 和数组的首地址
	return raw[:3]                           // 10000个字节只使用3个，其他的被浪费了
}

func goodGet() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 10000 10000 和数组的首地址
	res := make([]byte, 3)
	copy(res, raw[:3]) // 使用copy函数复制，raw可以被gc
	return res
}

func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])

	data2 := goodGet()
	fmt.Println(len(data2), cap(data2), &data2[0])

}
