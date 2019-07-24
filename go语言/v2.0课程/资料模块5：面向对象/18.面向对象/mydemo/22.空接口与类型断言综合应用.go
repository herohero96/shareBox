package main

import "fmt"

type Object struct {
	numA int
	numB int
}

type Add struct {
	Object
}

func (a * Add) GetResult(data ...interface{}) bool  {
	var b bool = true
	if len(data) > 2{
		fmt.Println("参数个数错误")
	}
	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数类型错误")
		b = false
	}
	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}

	a.numA = value
	a.numB = value1

	return b
}


type Sub struct {
	Object
}

func (s *Sub) SetData(data ...interface{}) bool  {
	var b bool = true
	if len(data) > 2 {
		fmt.Println("参数个数错误！！")
		b = false
	}
	value, ok := data[0].(int)
	if !ok {
		fmt.Println("第一个数类型错误")
		b = false
	}

	value1, ok1 := data[1].(int)
	if !ok1 {
		fmt.Println("第二个数据类型错误")
		b = false
	}

	s.numA = value
	s.numB = value1

	return b
}

func (s *Sub) GetResult() int {
	return s.numA - s.numB
}
type Resulter interface {
	GetResult() int
	SetData(data ...interface{}) bool
}