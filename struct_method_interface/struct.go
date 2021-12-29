package main

import "fmt"

func main()  { 
	tom := &Animal{
		name: "TOM",
	} 
	tom.hello("JERRY")
}

type Animal struct {
	name string;
	gender int;
}
func (animal *Animal) hello(name string) {
	fmt.Println("hello " + name + ", my name is " + animal.name + "!")
}