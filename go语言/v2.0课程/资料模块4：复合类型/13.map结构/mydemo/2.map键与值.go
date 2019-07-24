package main

import "fmt"

func main()  {

	m := map[int]string{1: "李四", 2: "王五"}

	value, ok := m[1]
	fmt.Println(value)
	fmt.Println(ok)

	for key, val := range m {
		fmt.Println(key)
		fmt.Println(val)
	}


	delete(m, 1)
	fmt.Println(m)


}