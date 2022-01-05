package main

import (
	"errors"
	"fmt"
	"os"
)

func main()  { 
	f, error := os.Open("aa.txt");
	fmt.Println(f)
	fmt.Println(error)
	if e := hello(""); e != nil {
		fmt.Println(e)
	}
	hello("chengpeng")
}

func hello(name string) error  {
	if len(name) == 0 {
		return errors.New("name is null!")
	}
	fmt.Println("hello " + name)
	return nil
}