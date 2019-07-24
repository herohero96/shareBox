package main

import "fmt"

func main() {
	var a int = 10
	var p *int

	p = &a
	fmt.Printf("%p\n", &a)
	fmt.Println(p)

	*p = 22
	fmt.Println(a)


}