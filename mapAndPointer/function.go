package main

import "fmt"

func main()  {
	num := 1;
	fmt.Println(num)
	add(num)
	fmt.Println(num)
	realAdd(&num)
	fmt.Println(num)
}

func add(num int)  {
	num += 1;
}

func realAdd(num *int)  {
	*num += 1;
}