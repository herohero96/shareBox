package User

import "fmt"

func Login(){
	fmt.Println("用户登录")
}
// 要想在别的文件中，调用该方法，方法名称的第一个字母必须大写。
func GetUser(){
	fmt.Println("获取用户信息")
}
