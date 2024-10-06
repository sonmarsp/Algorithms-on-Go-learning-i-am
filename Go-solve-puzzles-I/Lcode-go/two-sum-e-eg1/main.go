package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//if len(os.Args) == 1 {
	//	fmt.Println("Please give one or more floats.")
	//	os.Exit(1)
	//}

	arguments := os.Args

	//fmt.Println(arguments[1])
	//fmt.Println(arguments)

	strs := strings.Split(arguments[1], ",")

	//fmt.Println(strs)

	//переводим в интеджер из командной строки
	var ints []int
	for _, s := range strs {

		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}

	}

	fmt.Println(twoSum(ints, 10))

}

func twoSum(nums []int, target int) []int {

	c := len(nums)

	fmt.Println(c)
	fmt.Println(nums)

	var res []int

	for i := 0; i < c-1; i++ {

		for j := i + 1; j < c; j++ {

			fmt.Println(nums[i])
			fmt.Println(nums[j])
			fmt.Println("следующая пара")

			if (nums[i] + nums[j]) == target {

				fmt.Println("найдена пара")
				fmt.Println(nums[i])
				fmt.Println(nums[j])
				fmt.Println("найдена пара")

				res = append(res, i)
				res = append(res, j)

				return res

			}

		}

	}

	return res
}
