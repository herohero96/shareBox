package main

import "fmt"

func main()  {
	a := 10
	var p *int
	p = &a
	fmt.Println(p)
	var p1 *int
	p1 = new(int)
	*p1 = 22
	fmt.Println(p1)
}


