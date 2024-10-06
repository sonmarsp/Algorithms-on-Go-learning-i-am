package main

import "fmt"

func main() {

	fmt.Println("пример")

	arr := []int{1, 3, 5, 6}

	//value := 5
	value := 2

	result := binarySearch(arr, value)

	fmt.Println("targetT", result)

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

	return low // возвращаем -1, если значение не найдено
}

//Рабочий вариант но по скорости в 4 ms попадает в самые медленные
// а быстрый вариант просто перебор
func searchInsert(nums []int, target int) int {
	return findIndex(nums, len(nums), target)
}

func findIndex(arr []int, n, t int) int {
	for i := 0; i < n; i++ {
		if arr[i] == t {
			return i
		} else if arr[i] > t {
			return i
		}
	}
	return n
}
