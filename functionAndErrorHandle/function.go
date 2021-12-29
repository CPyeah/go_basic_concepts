package main

import "fmt"

func main()  { 
	r1 := add(1, 2)
	fmt.Printf(r1);
}

func add(a int, b int) int  {
	return a + b;
}

func div(a int, b int) (int, int)  {
	return a / b, a % b;
}