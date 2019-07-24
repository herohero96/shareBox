package main

import "fmt"

func main()  {
	type Student struct {
		id int
		name string
		age int
		addr string
	}

	var s Student = Student{101, "张三", 18, "北京"}

	fmt.Println(s)


	var s1 Student = Student{202, "李四", 22, "深圳"}
	var s2 Student = Student{id: 12, name: "王五"}

	s2.name = "二妮子"
	fmt.Println(s2)

	fmt.Println(s1)



}