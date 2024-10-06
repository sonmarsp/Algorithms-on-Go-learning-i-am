package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
6
2
3 6
3
1 2 4
3
3 6 1
3
15 35 21
4
35 10 35 14
5
1261 227821 143 4171 1941

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

	//fmt.Println(gcb(1, 4))

	scanner := bufio.NewScanner(os.Stdin)
	t := inputConsole(scanner)
	fmt.Println("Количество наборов")
	fmt.Println(t)

	i := 1
	for i <= t[0] {

		//fmt.Println("количество итераций")
		i++

		lp := inputConsole(scanner)
		l := lp[0]
		//fmt.Println("количество")
		//fmt.Println(l)

		ap := inputConsole(scanner)

		var res string = "No"

		for n := 1; n < l; n++ {

			//нужен еще цикл
		out:
			for r := l; r != 0; r-- {

				s := gcb(ap[r-1], ap[n-1])

				//fmt.Println(n - 1)
				//fmt.Println(n)
				//fmt.Println(ap[r-1])
				//fmt.Println(ap[n-1])

				//fmt.Println("Общий делитель")
				//fmt.Println(s)

				if l > s {
					res = "Yes"
					break out
				}

			} //end nested loop

		}

		fmt.Println(res)
	}

}

// Нужна функция для общего делителя
func gcb(a int, b int) int {

	//var c int
	for a != 0 && b != 0 {

		if a > b {
			a = a % b
		} else {
			b = b % a
		}

	}

	return a + b
}
