package main

import (
	"fmt"
)

func main() {
	var arrAge = [5]int{18, 20, 30, 15, 55}
	var arrName = [5]string{3: "gugu", 4: "xiaogu"}
	var arrCount = [4]int{500, 2: 100}
	var arrLazy = [...]int{1, 3, 4, 6, 8}
	var arrPack = [...]int{10, 5: 100}
	var arrRoom [20]int
	var arrBed = new([20]int)
	fmt.Println(arrAge)
	fmt.Println(arrName)
	fmt.Println(arrCount)
	fmt.Println(arrLazy)
	fmt.Println(arrPack)
	fmt.Println(arrRoom)
	fmt.Println(arrBed)

	var arr1 = new([5]int)
	arr := arr1
	arr1[2] = 100
	fmt.Println(arr[2], arr1[2])

	var arr2 [5]int
	arr3 := arr2
	arr2[2] = 100
	fmt.Println(arr2[2], arr3[2])

}
