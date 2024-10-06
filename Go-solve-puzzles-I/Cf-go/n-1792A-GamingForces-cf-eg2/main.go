package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Нужно обработать один ввод с переносом строк
3
4
1 2 1 2
3
2 4 2
5
1 2 3 4 5

*/

func main() {

	reader := bufio.NewReader(os.Stdin)
	t := readArrInt(reader)
	//defer reader.Flush()

	i := 1

	for i <= t[0] {
		i++

		np := readArrInt(reader)
		n := np[0]

		var cnt1 int = 0

		x := readArrInt(reader)
		for i := 0; i < n; i++ {

			if x[i] == 1 {
				cnt1++
			}
		}

		count := n - cnt1/2
		fmt.Println(count)
	}

}

// Yeap this is faster
func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

func readArrInt64(in *bufio.Reader) []int64 {
	numbs := readLineNumbs(in)
	arr := make([]int64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseInt(n, 10, 64)
		arr[i] = val
	}
	return arr
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

/*
Отработало засчитало задачу
но время 46мс а память 1100 кб
в некоторых тестах 3089 наборов данных
*/
