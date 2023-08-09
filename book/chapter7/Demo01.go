package main

import (
	"errors"
	"fmt"
	"strconv"
)

type error interface {
	Error() string
}

var atoiError = errors.New("time:invalid number")

func atoi(s string) (x int, err error) {
	neg := false
	if s != "" && (s[0] == '-' || s[0] == '+') {
		neg = s[0] == '-'
		s = s[1:]
	}
	q, err := strconv.Atoi(s)
	x = int(q)
	if err != nil {
		return 0, atoiError
	}
	if neg {
		x = -x
	}
	return x, nil

}

func main() {
	println(atoi("-100"))

	// 手动创建异常
	err := fmt.Errorf("aaaa ")
	if err != nil {
		println(err.Error())
	}
}
