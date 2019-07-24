package main

import "fmt"

func main()  {
/*	var Numbers[5] int = [5]int{1,2,3,4,5}
	fmt.Println(Numbers)
	fmt.Println(Numbers[0])*/

/*	Numbers := [5]int{1,2}
	fmt.Println(Numbers) //  默认值为0
	*/

/*	// 指定某个元素初始化
	Numbers := [5]int{2:5, 3:6}
	fmt.Println(Numbers)*/
/*
	Numbers := [...]int{7,8,9}
	fmt.Println(Numbers)
	fmt.Println(Numbers[0])

*/

/*	var Numbers [5] int
	Numbers[0]=1
	Numbers[1]=2
	fmt.Println(Numbers)

	for i :=0; i< len(Numbers); i++ {
		Numbers[i] = i
	}

	fmt.Println(Numbers)
*/
	var n1 [5]int = [5]int{1,2,3,4,5}
	n2 := [5]int{2:8, 3:9}
	n3 := [...]int{1,2,3,4}
	var n4 [5]int
	n4[0]=1

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)




}
