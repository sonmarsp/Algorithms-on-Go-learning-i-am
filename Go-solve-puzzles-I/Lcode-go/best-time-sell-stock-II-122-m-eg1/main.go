package main

import (
	"fmt"
)

func main() {

	//Input: prices = [7,1,5,3,6,4]
	//Output: 7

	prices := []int{7, 1, 5, 3, 6, 4}

	//prices := []int{2, 4, 1}
	fmt.Println(prices)

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	fmt.Println("maxProfit", maxProfit)

}
