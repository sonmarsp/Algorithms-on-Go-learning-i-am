package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1
5 2
10 20 5 15 25
*/

const MOD = 1_000_000_007

// inputConsole считывает строку чисел из консоли и возвращает срез целых чисел
func inputConsole(scanner *bufio.Scanner) []int64 {
	scanner.Scan()
	input := scanner.Text()
	ai := strings.Split(input, " ")

	var sint []int64
	for _, v := range ai {
		t, _ := strconv.ParseInt(v, 10, 64)
		sint = append(sint, t)
	}

	return sint
}

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	t := inputConsole(scanner)[0]

	for i := int64(0); i < t; i++ {
		nk := inputConsole(scanner)
		n := nk[0]
		k := nk[1]

		values := inputConsole(scanner)

		// Сортируем массив
		for j := int64(0); j < n-1; j++ {
			for l := int64(0); l < n-j-1; l++ {
				if values[l] < values[l+1] {
					values[l], values[l+1] = values[l+1], values[l]
				}
			}
		}

		aliceScore, bobScore := int64(0), int64(0)
		aliceTurn := true

		for j := int64(0); j < n; j++ {
			if aliceTurn {
				aliceScore += values[j]
				if j < k {
					continue
				}
			} else {
				bobScore += values[j]
			}
			aliceTurn = !aliceTurn
		}

		fmt.Fprintln(writer, aliceScore, bobScore)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(scanner, writer)
}
