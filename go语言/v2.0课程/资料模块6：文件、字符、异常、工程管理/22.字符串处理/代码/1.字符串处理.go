package main

import (
	"strings"
	"fmt"
)

// Contains : 判断一个字符串是否在另外一个字符串中。
// Join ：字符串连接
// Index ： 在一个字符串中查找某个字符串的位置
// Repeat ： 某个字符串重复多少次，返回的是重复后的字符串
// Replace ：在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
// Split ：把s字符串按照sep分割，返回slice（切片）
func main(){
   //var str string ="hellogo"
   //b:=strings.Contains(str,"goo") // 判断某个字符串是否在str中存在，如果存在返回true,否则返回false
   //fmt.Println(b)

   //join方法的使用
   //s:=[]string{"abc","hello","world"}
   //str:=strings.Join(s,"|") // 以"|"对切片中的内容进行分隔
   //fmt.Println(str)

   // Index
   //str:="abcHello"
   //n:=strings.Index(str,"Hello") // 判断"Hello"在str中出现的位置，注意位置从0开始计算
   //fmt.Println(n)

   // Repeat
   //str:=strings.Repeat("go",3) // 表示字符串"go"重复3次
   //fmt.Println(str)

   // Replace
   //str:="hello world"
   //fmt.Println(strings.Replace(str,"l","w",-1)) // 用新的字符串替换旧的字符串，第四个参数表示的替换的次数，如果小于0，表示全部替换

   // Split
   str:="hello@world@itcast"
   s:=strings.Split(str,"@")
   fmt.Println(s)
}
