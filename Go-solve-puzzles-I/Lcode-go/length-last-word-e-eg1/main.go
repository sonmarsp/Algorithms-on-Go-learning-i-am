package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	//Input: s = "Hello World"
	//Output: 5

	//str := "Hello World"
	//str := "aa"
	str := "a"

	/*
		for k, v := range str {
			fmt.Printf("Character %d: %c\n", k, v)
			if v == ' ' {
				fmt.Println("Found space!")
				break
			}
		}
	*/

	tStr := strings.TrimRightFunc(str, unicode.IsSpace)
	ls := len(tStr)
	fmt.Println(ls)

	r := 0
	for i := ls - 1; i >= 0; i-- {

		if tStr[i] == ' ' {
			r = ls - i - 1
			break
		} else {
			r = ls
			continue
		}
	}

	fmt.Println(r)
}
