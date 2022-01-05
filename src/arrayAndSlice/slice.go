package main 

import "fmt"

func main() {
	s1 := make([]int ,1)
	fmt.Println(s1)
	fmt.Println(s1[0])

	s2 := make([]int, 3,5);// [0, 0, 0, _, _]
	fmt.Println(s2)
	s2 = append(s2, 1);// [0, 0, 0, 1, _]
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))// 4, 5
	s2 = append(s2, 1, 2)
	fmt.Println(s2)// [0 0 0 1 1 2, _, _, _, _]
	fmt.Println(len(s2), cap(s2))// 6 10   容量变大了

	// 子切片
	s := s2[:2]
	fmt.Println(s)
	fmt.Println(s2[4:5])
	fmt.Println(s2)
	fmt.Println(s2[len(s2)-1:])

	// 合并切片
	s4 := s2[:1]
	s5 := s2[3:]
	s6 := append(s4, s5...)
	fmt.Println(s6)



}