package main

import "fmt"

func main()  {
	age := 12;
	if age < 18 {
		fmt.Println("kid")
	} else {
		fmt.Println("adult")
	}

	if age1 := 18; age1<18 {
		fmt.Println("kid")
	} else {
		fmt.Println("adult")
	}
	
}