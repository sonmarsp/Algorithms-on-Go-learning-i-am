package main

import (
	"fmt"
)

func main() {

	fmt.Println(addBinary("11", "1"))      // Output: "100"
	fmt.Println(addBinary("1010", "1011")) // Output: "10101"
}

func addBinary(a, b string) string {
	i, j := len(a)-1, len(b)-1
	c := 0
	res := ""

	for i >= 0 || j >= 0 || c > 0 {
		sum := c
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		res = string(sum%2+'0') + res
		c = sum / 2
	}
	return res
}
