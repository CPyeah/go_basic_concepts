package main

import "fmt"

func main()  { 
	fmt.Println(get(5))
}

func get(index int) (ret int)  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}