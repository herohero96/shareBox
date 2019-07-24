package main

import "fmt"

func main() {
/*	var s []int
	s = append(s, 1,2,3,4,5)
	s[3] = 88
	fmt.Println(s[3])
	fmt.Println(s)
*/
/*
	s := []int{8, 9, 7, 10, 12, 13}
	s = append(s, 99, 100)
	s[0] = 78
	fmt.Println(s[0])
	fmt.Println(s)*/

	s := make([]int, 3, 5)

	for i :=0; i< len(s); i++ {
		s[i] = i + 1
	}

	s = append(s, 80)
	s[3] = 90
	//s[4] = 100
	//s = append(s, 100)



	fmt.Println(s)




}
