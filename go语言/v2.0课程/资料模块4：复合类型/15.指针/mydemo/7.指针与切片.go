package main

import "fmt"

func main()  {
	s := []int{1,2,3,4,5}

	var p *[]int

	p = &s

	fmt.Println(*p)
	fmt.Println((*p)[0])


	for k, v := range *p {
		fmt.Println(k)
		fmt.Println(v)
	}




}