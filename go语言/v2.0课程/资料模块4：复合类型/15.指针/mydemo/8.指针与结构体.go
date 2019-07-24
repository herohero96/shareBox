package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

func main()  {
	stu := Student{12, "李四", 18}

	var p *Student
	p = &stu

	fmt.Println(p)

	fmt.Println(p.name)

}

