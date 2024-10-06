package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {

	// Замеряем время выполнения
	start := time.Now()

	nums1 := []int{0, 1, 2, 4, 5, 7}
	nums2 := []int{0, 2, 3, 4, 6, 8, 9}

	fmt.Println(ranges(nums1)) // Output: ["0->2", "4->5", "7"]
	fmt.Println(ranges(nums2)) // Output: ["0", "2->4", "6", "8->9"]

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
	// Получаем информацию об использовании памяти
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Выводим использование памяти
	fmt.Printf("Memory usage: Alloc = %v KB, TotalAlloc = %v KB, Sys = %v KB, NumGC = %v\n",
		memStats.Alloc/1024, memStats.TotalAlloc/1024, memStats.Sys/1024, memStats.NumGC)

}

func ranges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result []string
	start := nums[0]

	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			if i < len(nums) {
				start = nums[i]
			}
		}
	}

	return result
}

/*
Проход по массиву:

Индекс i = 1 (текущий элемент 1):

Проверка: 1 == 0 + 1 (элементы последовательные)
Индекс i = 2 (текущий элемент 2):

Проверка: 2 == 1 + 1 (элементы последовательные)
*/
