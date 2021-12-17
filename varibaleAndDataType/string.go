package main

import (
	"fmt"
	"reflect"
)

// 在 Go 语言中，字符串使用 UTF8 编码，
// UTF8 的好处在于，如果基本是英文，每个字符占 1 byte，
// 和 ASCII 编码是一样的，非常节省空间，如果是中文，一般占3字节。
// 包含中文的字符串的处理方式与纯 ASCII 码构成的字符串有点区别
func main() {
	str1 := "golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str1).Kind())
	fmt.Println(str2)
	fmt.Println(reflect.TypeOf(str2[3]))
	fmt.Println(len(str1))
	fmt.Println(len(str2))
	fmt.Println(str2[2])
}