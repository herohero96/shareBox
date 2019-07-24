package main

import (
	"strconv"
	"fmt"
)

// 把其他类型的转换为字符串。
// 把字符串转换为其他类型
func main() {
	// bool
	// int
	// 将bool类型转换成字符串类型
	//str:=strconv.FormatBool(true)
	//fmt.Println(str)

	// 将int类型转换成字符串类型
	//str:=strconv.Itoa(123)
	//fmt.Println(str)

	// 把字符串转换为其他类型
	// "true" --bool
	// "123" --int
	//b,err:=strconv.ParseBool("true")
	//if err!=nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println(b)
	//}

	// 字符串转成整型
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
