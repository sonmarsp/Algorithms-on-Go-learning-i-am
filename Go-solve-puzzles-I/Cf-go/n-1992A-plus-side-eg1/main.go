package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для вычисления максимального произведения
func maxP(a, b, c int) int {
	maxValue := a * b * c
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5-i; j++ {
			k := 5 - i - j
			newA := a + i
			newB := b + j
			newC := c + k
			product := newA * newB * newC
			if product > maxValue {
				maxValue = product
			}
		}
	}
	return maxValue
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Считываем количество тестов
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		nums := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		c, _ := strconv.Atoi(nums[2])

		result := maxP(a, b, c)
		fmt.Fprintln(writer, result)
	}
}
