package main

import "fmt"

func main()  {
	m1 := make(map[string]int)
	m2 := map[string]string {
		"Tom": "Male" ,
		"Marry" : "Female",
	}
	m1["Tome"] = 18;
	m2["Cat"] = "Animal"

	fmt.Println(m1)
	fmt.Println(m2)
}