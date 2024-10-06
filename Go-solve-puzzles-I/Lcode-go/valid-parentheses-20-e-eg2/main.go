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

Input: s = "{[]}" как такой вариант обработать
Output: true

Input: s = "[" как такой вариант обработать
Output: false

*/

func isValid(s string) bool {
	stack := []rune{}
	mb := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != mb[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	examples := []string{"()", "()[]{}", "(]", "{[]}", "[", "(("}

	for _, example := range examples {
		fmt.Printf("Input: %s, Output: %t\n", example, isValid(example))
	}
}
