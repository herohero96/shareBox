package main

import "fmt"

func main() {

	str := "helloword"

	m := make(map[byte]int)

	for i := 0; i< len(str); i++ {
		ch := str[i]
		m[ch] = m[ch] + 1
	}

	for key, value := range m {
		fmt.Printf("%c:%d\n", key, value)
	}





}
