package main

import "fmt"

func main()  {
	// 1：输入一个数字，判断是否为偶数，如果是输出"该数是偶数"，否则输出"该数是奇数"。
	 /*
     // （1）定义一个整型变量接收用户输入数字。
     var num int
     fmt.Println("请输入一个数字")
     fmt.Scan(&num)
     //  (2) 进行判断 能被2整除的数为偶数。
     if num%2==0 {
     	fmt.Println("输入的数字为偶数")
	 } else {
	 	fmt.Println("输入的数字为奇数")
	 }
     //  (3) 打印输出结果

  */
	// 2：输入公交卡当前的余额，只要超过2元，就可以上公交车；上车以后如果空座位的数量大于0，就可以坐下。

	 // (1) 定义变量接收用户输入的钱数
	 var money float64
	 fmt.Println("请输入钱数")
	 fmt.Scan(&money)
	 // (2) 对钱数进行判断
	 if money > 2{
		 // (3) 如果钱数大于2元，条件成立，判断座位数。
		 // (3.1) 要求用户输入座位数.
		 var count int
		 fmt.Println("请输入座位数：")
		 fmt.Scan(&count)
		 // (3.2) 对输入的座位数进行判断。
		 if count > 0{
		 	fmt.Println("你可以坐下！！")
		 } else {
		 	fmt.Println("你只能站着！！")
		 }
	 } else {
		 // (4) 如果钱数不满足，只能输出”不能上车”
		 fmt.Println("你不能上车！！")
	 }



}
