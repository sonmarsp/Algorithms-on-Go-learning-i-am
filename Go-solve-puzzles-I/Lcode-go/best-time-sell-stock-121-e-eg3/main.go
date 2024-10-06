package main

import (
	"fmt"
)

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{2, 4, 1}

	fmt.Println(prices)

	if len(prices) == 0 {
		fmt.Println("maxProfit", 0)
		return
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		fmt.Println("minPrice", minPrice)
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	fmt.Println("maxProfit", maxProfit)
}
