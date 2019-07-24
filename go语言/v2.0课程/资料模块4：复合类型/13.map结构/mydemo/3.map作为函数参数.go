package main

import "fmt"

func main()  {
	var m map[int]string = map[int]string{1:"张三", 2:"李四"}
	Printmap(m)

	DeleteMap(m)

	Printmap(m)

}

func Printmap(m map[int]string)  {
	for key, val := range m {
		fmt.Println(key)
		fmt.Println(val)
	}
}

func DeleteMap(m map[int]string)  {
	delete(m, 2)
}