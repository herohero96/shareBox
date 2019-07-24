package main

import "fmt"

func main()  {
	//var s[]int
	//s := []int{1,2,3}


	s := make([]int, 3, 5)
	fmt.Println(s)
	fmt.Println(cap(s))



}