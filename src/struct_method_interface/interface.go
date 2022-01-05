package main

import (
	"fmt"
	"reflect"
)

func main()  { 
	var tom Animal = &Cat{
		name: "TOM",
	}
	fmt.Println(tom.getName())
	fmt.Println(reflect.TypeOf(tom))

}

type Animal interface {
	getName() string;
}

type Cat struct {
	name string;
}
func (cat *Cat) getName() string {
	return cat.name;
}