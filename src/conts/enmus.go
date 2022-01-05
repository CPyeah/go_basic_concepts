package main

import "fmt"

const (
	a = 2 << iota
	b
	c
)

func main()  { 
		fmt.Println(a) 
		fmt.Println(b)
		fmt.Println(c)
}