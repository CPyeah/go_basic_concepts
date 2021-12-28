package main

import "fmt"

func main()  { 
	type Gender int8
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)
	
	gender := MALE
	
	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}

	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough
	case MALE:
		fmt.Println("male")
		fallthrough
	default:
		fmt.Println("unknown")
	}
	 
}