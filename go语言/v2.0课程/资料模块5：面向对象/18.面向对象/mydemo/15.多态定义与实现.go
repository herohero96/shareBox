/*package main

import "fmt"

type Personer interface {
	SayHello()
}

type Student struct {

}

func (s *Student)SayHello()  {
	fmt.Println("老师好")
}

type Teacher struct {

}

func (t *Teacher) SayHello()  {
	fmt.Println("学生好")
}

func  WhoSayHi(p Personer )  {
	p.SayHello()

}

func main()  {

	var stu Student
	WhoSayHi(&stu)

}
*/