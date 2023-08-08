package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 字符串拼接 性能差
	str := "bbbb" +
		"bb"
	println(str)

	// 性能一般
	str = fmt.Sprintf("%d:%s", 2023, "年")
	println(str)

	// 效率高，但是空间代价不小
	join := strings.Join([]string{"hello", "world"}, ",")
	println(join)

	// 理想
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(",")
	buffer.WriteString("world")

	println(buffer.String())

	// 理想，线程不安全
	var builder = strings.Builder{}
	builder.WriteString("hello")
	builder.WriteString(",")
	builder.WriteString("world")
	println(builder.String())

}
