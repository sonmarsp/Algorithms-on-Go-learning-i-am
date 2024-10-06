package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 2, 1, 1, 1, 2, 2, 3, 3}
	//nums := []int{1, 2, 3, 4}

	fmt.Println(nums)

	uMap := make(map[int]bool)
	uSlice := []int{}
	countNums := make(map[int]int)

	for _, num := range nums {
		if !uMap[num] {
			uMap[num] = true
			uSlice = append(uSlice, num)
			countNums[num]++
		} else {
			countNums[num]++
		}
	}

	var mk int
	var mv int

	for k, v := range countNums {
		if v > mv {
			mv = v
			mk = k
		}
	}

	fmt.Printf("Ключ с наибольшим значением: %d\n", mk)

	/*
		for i := 0; i < len(nums); i++ {

			for j := 0; j < len(uniqueSlice); j++ {
				if nums[i] == uniqueSlice[j] {
					countNums[uniqueSlice[j]]++
				}
			}
		}
	*/

	fmt.Println(uMap)
	fmt.Println(uSlice)
	fmt.Println(countNums)

}
