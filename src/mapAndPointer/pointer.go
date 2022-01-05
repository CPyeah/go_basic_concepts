package main

import "fmt"

func main()  {
	str := "hello"
	var p *string = &str // p是str的 指针
	*p = "hello world"
	fmt.Println(str)
}