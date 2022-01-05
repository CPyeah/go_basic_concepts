package main

import "fmt"

func main()  { 
	dog := struct{name string; age int}{name: "dd", age: 3}	 
	fmt.Println(dog)
	fmt.Println(dog.name)
}