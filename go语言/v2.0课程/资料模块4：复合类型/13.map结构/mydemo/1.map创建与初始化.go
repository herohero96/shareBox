package main

import "fmt"

func main()  {
/*	var n1 [5]int = [5]int{1,2,3,4,5}
	n2 := [5]int{2:8, 3:9}
	n3 := [...]int{1,2,3,4}
	var n4 [5]int
	n4[0]=1


	//var s[]int
	//s := []int{1,2,3}
	s := make([]int, 3, 5)
	*/

	var m map[int]string = map[int]string{1: "张三", 2: "李四", 3:"王五", 4:"itcast"}
	m1 := map[int]string{1: "张三", 2: "李四", 3: "王五", 4:"itcast"}
	m2 := make(map[string]int, 10)
	m2["张三"] = 12
	m2["李四"] = 15
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(len(m2))

}