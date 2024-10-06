package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

func solve(n int) int {
	dp := make([]int, n+2)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := i; j >= 0; j-- {
			dp[j+1] = (dp[j+1] + dp[j]) % MOD
		}
	}

	meow := 0
	for i := 1; i <= n+1; i++ {
		meow = (meow + dp[i]*i) % MOD
	}

	return meow
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Читаем количество тестов
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for ; t > 0; t-- {
		// Читаем n
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		// Вычисляем MEOW(a)
		result := solve(n)
		fmt.Fprintln(writer, result)
	}
}

//неверное решение
