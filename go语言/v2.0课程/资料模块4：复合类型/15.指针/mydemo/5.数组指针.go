package main

import "fmt"

func main()  {

	nums := [10]int{1,2,3,4,5,6,7,8,9,10}

	var p *[10]int

	p = &nums


	fmt.Println(p[1])
	fmt.Println((*p)[2])

	for i := 0; i < len(p) ; i ++ {
		fmt.Println(p[i])
	}

	UpdateArr(p)

	fmt.Println(*p)

}

func UpdateArr(p *[10]int)  {

	p[0] = 4321421

}
