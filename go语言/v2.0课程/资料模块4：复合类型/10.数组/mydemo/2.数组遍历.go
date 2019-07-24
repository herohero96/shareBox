package main

import "fmt"

func main()  {
	var Numbers [5]int = [5]int{1,2,3,4,5}

	for i := 0; i< len(Numbers); i++ {
		fmt.Println(Numbers[i])
	}

	for i, val := range Numbers {
		fmt.Printf("%d:%d\n", i, val)
	}





}