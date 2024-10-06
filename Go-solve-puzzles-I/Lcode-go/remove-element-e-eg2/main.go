package main

import (
	"fmt"
)

func main() {

	nums := []int{3, 2, 2, 3}
	needle := 3

	c := 0
	for _, v := range nums {
		if v != needle {
			nums[c] = v
			c++
		}
	}

	fmt.Println(nums)
	fmt.Println(c)

}
