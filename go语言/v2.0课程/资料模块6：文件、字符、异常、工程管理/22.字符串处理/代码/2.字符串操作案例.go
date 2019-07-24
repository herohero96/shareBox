package main

import (
	"fmt"
	"strings"
)

//让用户输入一个日期格式，如:2018-01-02, 输出日期为2008年1月2日
//让用户输入一句话,判断这句话中有没有“邪恶”,如果有“邪恶”就替换成“**”，然后输出.
func main() {
	// 1: 输入日期
	//fmt.Println("请输入日期，格式:年-月-日")
	//var str string
	//fmt.Scan(&str)
	//// 2: 按照"-"进行分隔
	//s:=strings.Split(str,"-")
	//// 3: 输出指定的格式
	//fmt.Println(s[0]+"年"+s[1]+"月"+s[2]+"日")

	// 第二道题
	// 1:定义变量存储用户输入的一句话
	fmt.Println("请输入一句话:")
	var str string
	fmt.Scan(&str)
	// 2:判断用户输入的内容中是否有“邪恶”
	if strings.Contains(str, "邪恶") {
		// 3: 如果有，则进行替换
		str = strings.Replace(str, "邪恶", "**", -1)
	}
	fmt.Println(str)

}
