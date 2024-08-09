package main

import (
	"fmt"
)

func main() {

	fmt.Println("бинарный поиск")
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	value := 7
	result := binarySearch(arr, value)

	if result != -1 {
		fmt.Printf("Элемент найден на индексе %d\n", result)
	} else {
		fmt.Println("Элемент не найден")
	}

}

func binarySearch(arr []int, value int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // возвращаем -1, если значение не найдено
}
