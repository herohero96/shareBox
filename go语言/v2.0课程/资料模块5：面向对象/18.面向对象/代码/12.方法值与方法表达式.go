package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) PrintInfo() {
	fmt.Println(*p)
}

func main() {
	per := Person{"张三", 18}
	per.PrintInfo()

	//方法值。
	//f := per.PrintInfo
	//fmt.Printf("%T\n",f)
	//f()
	//方法表达式
	f:=(*Person).PrintInfo
	f(&per)

}
