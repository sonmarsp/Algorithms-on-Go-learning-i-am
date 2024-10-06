package main

import (
	"fmt"
)

func main() {

	//nums := []int{2, 2, 1, 1, 1, 2, 2, 3, 3}
	nums := []int{1, 1, 1, 2, 2, 3}

	mMap := make(map[int]int)
	c := 0

	/*
		for i := 0; i < len(nums); i++ {
			if mMap[nums[i]] < 2 {
				mMap[nums[i]] = mMap[nums[i]] + 1
				nums[c] = nums[i]
				c++
			}
		}
	*/

	for _, num := range nums {
		if mMap[num] < 2 {
			mMap[num] = mMap[num] + 1
			nums[c] = num
			c++
		}
	}

	fmt.Println(nums)
	//fmt.Println(c)
}
