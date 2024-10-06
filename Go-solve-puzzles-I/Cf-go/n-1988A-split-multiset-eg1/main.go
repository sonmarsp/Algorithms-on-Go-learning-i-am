package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(n int, k int) int {
	oper := 0

	for n > 1 {
		n = n - k + 1
		oper++
	}
	return oper

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		params := scanner.Text()
		var n, k int
		fmt.Sscanf(params, "%d %d", &n, &k)

		result := solve(n, k)
		fmt.Fprintln(writer, result)
	}
}

//верное формула была
/*
4
1 5
5 2
6 3
16 4
----- правильный вывод
0
4
3
5
--------

1
5 2
*/
