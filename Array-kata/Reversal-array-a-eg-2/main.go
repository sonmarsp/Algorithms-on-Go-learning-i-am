package main

import (
	"fmt"
)

func main() {

	//arr := []int{3, 5, 6, 1, 7, 11, 14, 8, 18, 17, 43, 34}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	//arr := []int{1, 2, 3}
	fmt.Println(arr)

	reverse(arr, 0, len(arr)-1)
	fmt.Println(arr)

}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
