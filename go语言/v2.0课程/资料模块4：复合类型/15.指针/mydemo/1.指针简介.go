package main

import "fmt"

/*type Student struct {
	id int
	name string
}
*/
func main()  {
	var num int = 10
	Test(num)
	fmt.Println(num)

	var m []int = []int{1,2,3,4}
	Test2(m)
	fmt.Println(m)

	arr := [...]int{1,2,3,4}
	Test3(arr)
	fmt.Println(arr)

	var map1 map[int]string = map[int]string{1: "李四", 2: "王五"}

	Test4(map1)
	fmt.Println(map1)

	var stu Student = Student{12, "爸爸"}
	Test5(stu)
	fmt.Println(stu)
}

func Test(n int)  {
	n = 30
	fmt.Println(n)
}

func Test2(m []int )  {
	m[0] = 90
	fmt.Println(m)
}
func Test3(arr [4]int)  {
	arr[0] = 900
	fmt.Println(arr)
}

func Test4( m map[int]string )  {
	m[1] = "张三"
	fmt.Println(m)
}

func Test5(stu  Student )  {
	stu.name = "妈妈"
	fmt.Println(stu)
}

