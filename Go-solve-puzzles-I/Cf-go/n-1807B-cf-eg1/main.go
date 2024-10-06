package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
5
4 7 7 2 8
*/

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

		np := inputConsole(scanner)
		n := np[0]

		//fmt.Println(n)

		//var cnt1 int = 0

		x := inputConsole(scanner)
		//fmt.Println(x)

		var c int = 0

		for i := 0; i < n; i++ {

			if x[i]%2 == 0 {
				c++
			}
		}

		//fmt.Println("count del")
		//fmt.Println(c)

		if c >= (n+1)/2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}
