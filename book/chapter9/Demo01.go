package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string "姓名" // 结构体标签
	age  int    "年龄"
	Room int    `json:"Roomid"`
}
type Human struct {
	name string
}

type Student2 struct {
	Human // 匿名（内嵌）字段
	int
}

type Person1 struct {
	Human // 内嵌
}

type Person2 struct {
	*Human // 内嵌 与Person1 不同
}
type Person3 struct {
	human Human // 聚合
}

type Reader interface {
	Reader()
}

type Teacher interface {
	Reader
	//Writer
	//Human 接口中只能嵌套接口 不能嵌入其他类型
}

func main() {
	s := Student{"gugu", 28, 1101}
	fmt.Println(reflect.TypeOf(s).Field(0).Tag)
	fmt.Println(reflect.TypeOf(s).Field(1).Tag)
	fmt.Println(reflect.TypeOf(s).Field(2).Tag)

}
