package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func main()  {
	stu1 := Student{"里斯", 100}
	stu2 := Student{"张三", 20}
	stu3 := Student{"王五", 80}


	p := []Student{stu1, stu2, stu3}

	for k,  v := range p {
		v.Age = v.Age + 1000
		p[k].Age = v.Age
	}

	fmt.Println(p)




}