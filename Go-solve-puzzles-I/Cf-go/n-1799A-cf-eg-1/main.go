package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
10
1 1
2
3 2
5 4
4 5
5 9 9 5 7
5 5
6 7 8 9 10
3 4
4 4 4 4
4 4
5 5 6 6
3 5
4 5 5 5 4
4 20
5 5 24 24 24 5 6 7 8 9 10 12 13 14 15 16 17 18 19 20
5 7
7 8 7 11 7 12 10
6 7
8 11 7 8 8 8 12
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
		l1 := lp[0] //размер поля прямой эфир
		l2 := lp[1]

		//fmt.Println("количество")
		//fmt.Println(l1)
		//fmt.Println(l2)

		var life []int
		for l := 1; l <= l1; l++ {

			fmt.Println(l)
			life = append(life, l)
		}
		//life[2] = 5
		ar := inputConsole(scanner)
		fmt.Println(ar)

		//вторая итерация замена от количества ходов
		//идея создать один массив
		/*идея создать один массив
		          1 2 3 4 5
				  5 1 2 3 4
				  4 5 1 2 3
				  и брать цифры из прямого эфира
				  посчитаь через сколько шагов их индекс меньше диапозона прямого эфира
		Нужно научиться перемешать в Go по массиву.
		*/

		for m := 1; m <= l2; m++ {

			fmt.Println("Количество действий")
			fmt.Println(m)

			life[m] = ar[m]

		}

		fmt.Println("life")
		fmt.Println(life)

	} //end main loop

}
