package main

import "fmt"

func main()  { 
	r1 := add(1, 2)
	fmt.Println(r1)
	r2, mod := div(10, 3)
	fmt.Println(r2, mod)
}

func add(a int, b int) int  {
	return a + b;
}

func div(a int, b int) (int, int)  {
	return a / b, a % b;
}