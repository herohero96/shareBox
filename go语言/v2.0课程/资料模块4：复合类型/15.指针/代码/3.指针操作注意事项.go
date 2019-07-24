package main

import "fmt"

/*
func main() {
	var p *int
	fmt.Println(p)
	//*p = 78
	//p = new(int)
	//*p = 67
	//fmt.Println(*p)
}*/

func main()  {
	var a int = 10
	var p *int
	p = &a

	fmt.Println(p)

	var p1 *int
	p1 = new(int)
	*p1 = 88
	fmt.Println(*p1)



}



