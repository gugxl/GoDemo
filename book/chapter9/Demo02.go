package main

import "fmt"

type Writer interface {
	Writer()
}
type Author struct {
	name string
	Writer
}
type Other struct {
	i int
}

func (a Author) Write() {
	fmt.Println(a.name, " Write。")
}

func (o Other) Writer() {
	fmt.Println(" Write。")
}

func main() {
	Ao := Author{"Other", Other{12}}
	Ao.Write()

	Au := Author{name: "Hawing"}
	Au.Write()
}
