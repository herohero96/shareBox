package main

import (
	"fmt"
)

func main()  {


	n := make([]int, 3, 5)
	InitData(n)

	sum := SumAdd(n)

	fmt.Println("sum:",sum)


}

func InitData( num []int)  {
	for i := 0; i< len(num) ; i++ {
		fmt.Printf("请输入第%d个数\n", i+1)
		fmt.Scan(&num[i])

	}
}

func SumAdd(num []int) int  {
	var sum int
	for i := 0; i < len(num) ; i++ {
		sum += num[i]
	}
	return sum
}