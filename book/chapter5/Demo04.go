package main

import "fmt"

func main() {
	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	s3 := append(s2, s0...)
	s4 := append(s3[3:6], s3[2:]...)
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("s3=", s3)
	fmt.Println(s4)

}
