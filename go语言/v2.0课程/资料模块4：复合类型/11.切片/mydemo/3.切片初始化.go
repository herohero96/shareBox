package main

import "fmt"

func main()  {
/*	var s []int
	s = append(s, 1,2,3,4,5,6)
	fmt.Println(s)

	s[3] = 8
	fmt.Println(s)
*/

/*	s := []int{1,2}
	s = append(s, 88, 99)*/

	s := make([]int, 3, 10)
	for i := 0; i< len(s) ; i++  {
		s[i] = i+1
	}
	s = append(s, 80)

	fmt.Println(s)



}