package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	// store 存储数据
	m.Store("name", "gugu")
	m.Store("gender", "man")

	//	lockOrstore
	//	如果key不存在，就存储key，value，返回false和输入的value
	v, ok := m.LoadOrStore("name1", "gxl")
	fmt.Println(ok, v) // false gxl

	//如果key存在。返回true和对应的value，不会对原来的value修改
	v, ok = m.LoadOrStore("name1", "gxl22")
	fmt.Println(ok, v) // true gxl

	// Load读取数据
	v, ok = m.Load("name1")
	if ok {
		fmt.Println("key存在，值是：", v)
	} else {
		fmt.Println("key 不存在")
	}

	// 遍历
	f := func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	}
	m.Range(f)

	// Delete 删除数据
	m.Delete("name1")
	fmt.Println(m.Load("name1"))
}
