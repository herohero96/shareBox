package main

import "fmt"

func main()  {
 var str string = "helloWorld"

 m := make(map[byte]int, 10)

	for i := 0; i < len(str); i ++ {
		ch := str[i]
		m[ch] = m[ch] + 1
	}

	for key, value := range m {
		fmt.Printf("%c:%d\n", key, value)
	}

}