package main

import "fmt"

func main() {
	var arr1 [5]int;
	fmt.Println(arr1)

	var arr2 = [3]int{1,2,3}
	fmt.Println(arr2)

	arr3 := [5]int{2,3,4,6,7}
	for i := 0; i < len(arr3); i++ {
		arr3[i] = arr3[i] *100
	}
	fmt.Println(arr3)
	

}