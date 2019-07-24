package main

import "fmt"

func main()  {
	Numbers := [5]int{2,5,7,8,9}

	var max int = Numbers[0]
	var min int = Numbers[0]
	var sum int
	var average float64
	// 最大值

	for _, val := range Numbers{
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
		sum += val
	}
	average = float64(sum)/float64(len(Numbers))



	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(sum)
	fmt.Println(average)





}