package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//const N = 5010
//const MOD = 1_000_000_007

const N = 6
const MOD = 1_000_000_007

var c [N][N]int

func initCombinations() {
	for i := 0; i < N; i++ {
		c[i][0] = 1
		c[i][i] = 1
		for j := 1; j < i; j++ {
			c[i][j] = (c[i-1][j-1] + c[i-1][j]) % MOD
		}
	}

}

func solve(n int) int {
	ans := 0
	for i := 1; i <= 2*n+1; i++ {
		for x := 0; x <= min(i-1, n); x++ {
			y := i - 2*x - 1
			if !(0 <= y && y <= max(0, n-i)) {
				continue
			}
			ans = (ans + c[max(0, n-i)][y]*c[min(i-1, n)][x]%MOD*i%MOD) % MOD
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	initCombinations()

	fmt.Println(c)

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for t := 0; t < T; t++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		result := solve(n)
		fmt.Fprintln(writer, result)
	}
}

//не верное решение
/*
1
2
*/
