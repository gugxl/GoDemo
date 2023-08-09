package main

import "fmt"

func main() {
	var a, b int = 20, 30
	var ptra *int
	var ptrb *int = &b
	ptra = &a

	fmt.Printf("&a %x \n", &a)
	fmt.Printf("&b %x \n", &b)

	fmt.Printf("ptra %x \n", ptra)
	fmt.Printf("ptrb %x \n", ptrb)

	fmt.Printf("ptra 的值%d \n", *ptra)
	fmt.Printf("ptrb  的值 %d \n", *ptrb)

}
