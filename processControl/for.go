package main

import "fmt"

func main()  { 
	sum := 0;
	for i := 0; i <= 100; i++ {
		sum += i;
	}	 
	fmt.Println(sum)

	nums := []int{1, 2,3,4,5,6,7,8,9,10}
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])		
	}
	fmt.Println("")
	for i, num:=range nums {
		fmt.Println(i, num)
	}

	map1 := map[string]string {
		"name": "chengpeng",
		"age": "18",
	}
	for k,v := range map1{
		fmt.Println(k, v)
	}
}