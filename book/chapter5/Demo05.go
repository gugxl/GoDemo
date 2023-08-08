package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1)
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2)
	for i := range s2 {
		s2[i] += 20
	}
	// s2和s1指向同一个数组，修改s2会影响s1
	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 5) // 这里扩容会指向新数组
	for i := range s2 {
		s2[i] += 20
	}
	// s1还是老数据， s2扩容后不会影响s1
	fmt.Println(s1)
	fmt.Println(s2)
}
