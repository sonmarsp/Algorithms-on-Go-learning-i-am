package main

import (
	"bufio"
	"fmt"
	"os"
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

	i := 1

	for i <= t[0] {
		i++

		q := inputConsole(scanner)
		//n := np[0]
		//fmt.Println(q)
		if (q[0] + q[1]) == q[2] {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}

	}

}
