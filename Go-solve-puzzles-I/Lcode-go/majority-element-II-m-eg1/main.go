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

	findElem := []int{}

	for k, v := range countNums {
		if v > len(nums)/3 {
			findElem = append(findElem, k)
		}
	}

	fmt.Println(uMap)
	fmt.Println(uSlice)
	fmt.Println(countNums)
	fmt.Println(findElem)

}

func uniqueMapEg(nums []int) []int {
	uniqueMap := make(map[int]bool)
	uniqueSlice := []int{}

	for _, num := range nums {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}

	return uniqueSlice
}
