package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func inputConsole(scanner *bufio.Scanner) []int {

	scanner.Scan()
	input := scanner.Text()
	ai := strings.Split(input, " ")

	var sint []int
	for _, v := range ai {

		t, _ := strconv.Atoi(v)
		sint = append(sint, t)
	}

	return sint
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := inputConsole(scanner)
	fmt.Println(t)

	i := 1

	for i <= t[0] {
		i++

		np := inputConsole(scanner)
		//n := np[0]

		fmt.Println("np", np)

		x := inputConsole(scanner)
		fmt.Println("x", x)

		solve(x)
		fmt.Println("mutate", x)

	}

}

// Проверям на нули
func allZeros(a []int) bool {
	for _, x := range a {
		if x != 0 {
			return false
		}
	}
	return true
}

// Находим медиану
func findMedian(a []int) int {

	sort.Ints(a)
	n := len(a)
	if n%2 == 1 {

		return a[n/2]
	}
	return (a[n/2-1] + a[n/2]) / 2
}

func findMean(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum / len(a)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(a []int) {

	op := 40

	for i := 0; i <= op; i++ {
		//median := findMedian(a)
		median := findMean(a)
		//fmt.Println("mediana", findMedian(a))
		fmt.Println("arr", a)

		for j := 0; j < len(a); j++ {
			a[j] = abs(a[j] - median)

			if allZeros(a) {
				fmt.Println("все нули")
				continue
			}
		}
	}

}

/*
5
1
5
2
0 0
3
4 6 8
4
80 40 20 10
5
1 2 3 4 5

1
4
80 40 20 10

1
5
1 2 3 4 5

*/
