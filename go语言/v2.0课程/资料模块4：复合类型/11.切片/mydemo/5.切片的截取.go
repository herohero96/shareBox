package main

import "fmt"

func main()  {
	s := []int{1,2,3,4,5}
/*
	s1 := s[1:3:5]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap( s1))
	s[1] = 100
	fmt.Println(s1)

	s1[0] = 88
	fmt.Println(s)
*/


	s1 := s[1:3]
	fmt.Println(s1)
	fmt.Println(cap(s1))


}