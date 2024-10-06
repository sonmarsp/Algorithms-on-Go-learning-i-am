package main

import (
	"fmt"
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

	//var c int = m + n

	//fmt.Println(c)

	nums1 = nums1[:m]
	//fmt.Println("nums1")
	//fmt.Println(nums1)

	for i := 0; i < n; i++ {

		nums1 = append(nums1, nums2[i])
	}

	//nums1 = append(nums1[:m], nums2...)
	fmt.Println(nums1)

	//здесь вызвать сортировку
	nums1 = bubbleSort(nums1)
	//bubbleSort(nums1)

	// Sort nums1
	//nums1 = sort.Ints(nums1)

	//fmt.Println(nums1)
	//fmt.Println(bs)
	return nums1
}

func bubbleSort(arr []int) []int {

	if len(arr) > 1 {
		for i := 0; i < len(arr)-1; i++ {
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}

		}
	}

	return arr
}
