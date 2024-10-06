package main

import (
	"fmt"
)

/**
Нужно обработать один ввод с переносом строк
3
3
6
7

*/

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		result := genArray(n)
		for _, val := range result {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func genArray(n int) []int {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i * 105 / gcd(i, 105)
	}
	return arr
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
