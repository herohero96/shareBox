package main

import "fmt"

func main() {
	var a int = 10
	var p *int
	p = &a
	//fmt.Printf("%p\n", &a)
	//fmt.Printf("%p", p)
	fmt.Println(*p)
	*p=222
	fmt.Println("a=",a)
}
