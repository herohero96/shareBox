package main

import "fmt"

func main()  {
	var n1 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(n1)
	n2 := [...][2]int{{1,2},{4,5}}
	fmt.Println(n2)

	n3 := [5][3]int{1:{1,2,3}, 4:{7,8,9}}
	fmt.Println(n3)

	var n4 [2][3]int = [2][3]int{{1}}
	fmt.Println(n4)

}