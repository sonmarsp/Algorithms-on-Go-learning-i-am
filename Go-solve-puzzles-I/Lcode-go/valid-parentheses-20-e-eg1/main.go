package main

import "fmt"

/*
Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Input: s = "{[]}" на этом не верно
Output: true
*/

func main() {

	s := make(map[int]int)
	s[40] = 41
	s[91] = 93
	s[123] = 125

	fmt.Println(s)

	//str := "()"
	str := "()[]{}"
	fmt.Println(str)

	for i := 0; i < len(str)-1; i++ {
		fmt.Println("str[i]-- ", str[i])
		fmt.Println("s[i]-- ", s[int(str[i])])
		fmt.Println("i ", i)

		if s[int(str[i])] == int(str[i+1]) && i%2 == 0 {
			fmt.Println("true")
		} else if i%2 == 0 {
			fmt.Println("false")
		}
	}

}
