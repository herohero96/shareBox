package main

import "fmt"

func main()  {
	var p [2]*int
	var i int = 10
	var j int = 20

	p[0] = &i
	p[1] = &j

	fmt.Println(p)

	fmt.Println(*p[1])

	for _, v := range p {
		fmt.Println(*v)
	}

}