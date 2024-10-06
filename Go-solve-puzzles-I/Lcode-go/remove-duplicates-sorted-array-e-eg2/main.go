package main

import (
	"fmt"
)

func main() {

	//вспоминаю про абстрактный тип данных
	// нужно основные операции с массивами попробовать смещение, разворот
	// в данном случае смещение в конец
	//nums := []int{1, 1, 2}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println("nums--")
	fmt.Println(nums)

	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[l] = nums[i]
			l++
		}
		prev = nums[i]
	}

	fmt.Println("change nums--")
	fmt.Println(nums)
}

//По условию задачи над тем же вводом
