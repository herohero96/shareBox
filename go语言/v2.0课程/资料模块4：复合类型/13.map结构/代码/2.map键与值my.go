package main

import "fmt"

func main() {
	var m map[int]string = map[int]string{1: "王五", 2: "李四"}
	fmt.Println(m)

	value, ok := m[7]
	fmt.Println(value)
	fmt.Println(ok)

	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value)
	}


	delete(m, 2)
	fmt.Println(m)



}
