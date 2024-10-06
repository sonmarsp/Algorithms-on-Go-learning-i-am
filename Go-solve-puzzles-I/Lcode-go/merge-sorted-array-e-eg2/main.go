package main

import (
	"fmt"
	"sort"
)

func main() {

	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//nums2 := []int{2, 5, 6}

	nums1 := []int{0}
	nums2 := []int{1}

	m := 0
	n := 1

	nums1 = merge(nums1, m, nums2, n)

	fmt.Println(nums1)
	//fmt.Println(nums2)

}

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	for n > 0 {
		nums1[m] = nums2[n-1]
		m++
		n--
	}
	sort.Ints(nums1)

	return nums1

}
