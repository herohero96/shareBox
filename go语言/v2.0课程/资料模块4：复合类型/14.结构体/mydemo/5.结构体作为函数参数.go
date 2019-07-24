package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main()  {
	stu :=  make([]Student, 2, 10)

	InitData(stu)
	GetMax(stu)

}

func InitData(stu []Student)  {
	for i :=0; i <len(stu) ; i++  {
		fmt.Printf("请输入第%d个学生的详细信息\n", i + 1)
		fmt.Scan(&stu[i].id, &stu[i].name, &stu[i].age, &stu[i].addr)
	}
}

func GetMax(stu []Student)  {
	var max int = stu[0].age
	var maxIndex int
	for i := 0; i < len(stu) ; i++  {
		if stu[i].age > max {
			max = stu[i].age
			maxIndex = i
		}
	}
	fmt.Println(stu[maxIndex])
}