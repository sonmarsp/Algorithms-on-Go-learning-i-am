package main

import (
	"fmt"
)

func main() {

	//arr := []int{3, 5, 6, 1, 7, 11, 14, 8, 18, 17, 43, 34}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	//k := 3
	//rotateLeft(arr, k)
	fmt.Println(arr)

	//reverse(arr, 0, len(arr)-1)
	rotateLeft(arr, 3)
	fmt.Println(arr)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func rotateLeft(arr []int, k int) {
	n := len(arr)
	k = k % n // Если k больше длины массива
	reverse(arr, 0, n-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)

}
