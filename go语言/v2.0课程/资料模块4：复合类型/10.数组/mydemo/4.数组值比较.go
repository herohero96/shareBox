package main

import "fmt"

func compare(n1 [5]int, n2 [5]int)  {
	if( len(n1) == len(n2)){
		for i:=0; i<len(n1); i++ {
			if( n1[i] != n2[i] ){
				break
			}
			if (i == len(n1)-1){
				fmt.Println("这2个数组相等")
			}
	}
}



}



func main()  {
	n1 := [5]int{1,2,3,4,}
	n2 := [5]int{1,2,3,4,}
	compare(n1, n2)


}
