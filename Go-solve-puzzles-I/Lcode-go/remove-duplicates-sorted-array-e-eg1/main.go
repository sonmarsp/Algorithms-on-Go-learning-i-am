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

	var tmp1 []int
	var tmp2 []int
	c := 0

	for i := 0; i < len(nums)-1; i++ {

		//fmt.Println(i)

		if nums[i] != nums[i+1] {

			tmp1 = append(tmp1, nums[i])

			fmt.Println(nums[i])
			fmt.Println(nums[i+1])
			c++
		} else {
			tmp2 = append(tmp2, nums[i])
		}

		if i == len(nums)-2 && nums[i] != nums[i+1] {
			tmp1 = append(tmp1, nums[i+1])
			c++
		}

		//иначе в конец массива попробовать
		//fmt.Println(nums)
	}

	fmt.Println(tmp1)
	fmt.Println(tmp2)

	tmp1 = append(tmp1, tmp2...)

	fmt.Println(tmp1)
	fmt.Println(c)
}

//По условию задачи над тем же вводом
