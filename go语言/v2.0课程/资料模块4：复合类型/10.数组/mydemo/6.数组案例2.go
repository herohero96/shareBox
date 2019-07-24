package main

import "fmt"

func main()  {

	names := [...]string{"马龙", "迈克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩特"}

	maxName := names[0]

	for _, val := range names {
		if len(val) > len(maxName) {
			 maxName = val
		}
	}

	fmt.Println(" 最大字符串", maxName)



}
