package main

import "fmt"

type People struct {
	Age    int
	gender string
	Name   string
}

type OtherPeople struct {
	People
}

func (p People) PeInfo() {
	fmt.Println("people", p.Name, ":", p.Age, "岁,性别", p.gender)
}
func main() {
	//OtherPeople.PeInfo()
}
