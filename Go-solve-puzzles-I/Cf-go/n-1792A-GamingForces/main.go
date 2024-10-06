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

в vscode powershel перенос строк "строка1 `n  строка `n"
*/


/*Не работает на codeforces без указателей */
func inputConsole() []int {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	ai := strings.Split(input, " ")

	//fmt.Println("inputConsole")
	//fmt.Println("Вывод")
	//fmt.Println(input)
	//fmt.Println(ai)

	var sint []int
	for _, v := range ai {
		//fmt.Println(i)
		//fmt.Println(v)
		t, _ := strconv.Atoi(v)
		sint = append(sint, t)
	}

	return sint
}

func main() {

	t := inputConsole()
	//fmt.Println(t)

	i := 1

	for i <= t[0] {
		//fmt.Println(i)
		i++

		np := inputConsole()
		n := np[0]
		//fmt.Println(n)

		var cnt1 int = 0

		x := inputConsole()
		for i := 0; i < n; i++ {

			//fmt.Println("Вывод данных")
			//fmt.Println(x[i])

			if x[i] == 1 {
				cnt1++
			}
		}

		count := n - cnt1/2
		//fmt.Println("count")
		fmt.Println(count)
	}

}
