package main

import (
	"fmt"
)

func main() {
	//start1 := time.Now()

	arr := []int{3, 5, 6, 1, 7, 11, 14, 8, 18, 17, 43}

	//fmt.Println(BubbleSortEg1(arr))

	fmt.Println(BubbleSortEg2(arr))

	//duration1 := time.Since(start1)
	//fmt.Println(duration1)
}

func BubbleSortEg1(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}

	}

	return arr
}

func BubbleSortEg2(input []int) []int {
	n := len(input)
	swapped := true

	for swapped {

		swapped = false

		for i := 0; i < n-1; i++ {

			if input[i] > input[i+1] {
				fmt.Println("Swapping")

				input[i], input[i+1] = input[i+1], input[i]

				swapped = true
			}
		}
	}

	return input
}
