package main

import "fmt"

func main() {

	var m map[int]string = map[int]string{1: "王五", 2: "李四"}
	// fmt.Println(m[2])

	value, ok := m[6]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println(value)
		fmt.Println("不存在")
	}

	/*
	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value)
	}
	*/
	delete(m,2)
	fmt.Println(m)
}
