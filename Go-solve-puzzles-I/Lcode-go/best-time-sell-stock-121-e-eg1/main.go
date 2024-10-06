package main

import (
	"fmt"
)

func main() {

	//как я решу эту задачу графически в уме?
	//получить максимальную прибыль нужен Output: 5
	//prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{2, 4, 1}

	fmt.Println(prices)

	l := len(prices)
	//fmt.Println(l)

	min := prices[0]
	key := 0

	//поиск минимального элемента в массиве и его индекс
	for i := 0; i < l; i++ {
		fmt.Println(i)
		if prices[i] < min {
			min = prices[i]
			key = i
		}
	}

	fmt.Println("min")
	fmt.Println(min)
	fmt.Println("key")

	maxProfit := 0
	for i := key; i < l; i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
	}

	fmt.Println("maxProfit", maxProfit)

}
