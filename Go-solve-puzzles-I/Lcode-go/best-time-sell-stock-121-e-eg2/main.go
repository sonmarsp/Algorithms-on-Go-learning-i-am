package main

import (
	"fmt"
)

func main() {

	//как я решу эту задачу графически в уме?
	//получить максимальную прибыль нужен Output: 5
	prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{2, 4, 1}

	fmt.Println(prices)

	l := len(prices)
	//fmt.Println(l)

	maxProfit := 0
	for i := 0; i < l; i++ {
		fmt.Println(prices[i])
		for j := i + 1; j < l; j++ {
			fmt.Println(prices[j])
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}

	}

	fmt.Println("maxProfit", maxProfit)

	// выбивает по времени на большом массиве
}
