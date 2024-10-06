package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("пример")
	//1 шаг развернуть строку в обратном порядке

	x := -123

	s := strconv.Itoa(x)
	fmt.Println("s", s)
	n := len(s)

	res := true
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			res = false
		}
	}

	fmt.Println("res", res)
}
