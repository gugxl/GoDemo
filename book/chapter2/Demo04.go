package main

import "unicode/utf8"

func main() {
	const (
		a = iota
		b
		c
	)

	var d = 10
	println(a)
	println(b)
	println(c)
	println(d)

	println(1i * 1i)

	s := "其实就是rune"
	println(len(s))
	println(utf8.RuneCountInString(s))

}
