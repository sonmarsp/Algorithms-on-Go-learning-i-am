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

		var cnt1 int = 0

		x := inputConsole(scanner)
		for i := 0; i < n; i++ {

			if x[i] == 1 {
				cnt1++
			}
		}

		count := n - cnt1/2
		fmt.Println(count)
	}

}

/*
Отработало засчитало задачу
но время 46мс а память 900 кб
в некоторых тестах 3089 наборов данных
*/
