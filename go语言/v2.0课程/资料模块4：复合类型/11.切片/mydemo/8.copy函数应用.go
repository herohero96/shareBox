package main

import "fmt"

func main()  {
	s1 := []int{1,2}
	s2 := []int{3,4,5,6}
	//copy(s1, s2)
	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)




}