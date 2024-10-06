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

func maxElementForward(arr []int, n int) int {
	//n := len(arr)
	maxV := arr[0]

	for i := 0; i < n; i += 2 {
		if arr[i] > maxV {
			maxV = arr[i]
		}
	}

	return maxV
}

// Функция для нахождения максимального элемента с шагом 2 в обратном порядке
func maxElementBackward(arr []int, n int) int {

	maxV := arr[n-1]

	// Проходим по массиву в обратном порядке с шагом 2
	for i := n - 1; i >= 0; i -= 2 {
		if arr[i] > maxV {
			maxV = arr[i]
		}
	}

	return maxV
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := inputConsole(scanner)
	fmt.Println(t)

	i := 1

	for i <= t[0] {
		i++

		np := inputConsole(scanner)
		n := np[0]

		fmt.Println("np", np)

		x := inputConsole(scanner)
		fmt.Println("x", x)
		fmt.Println("Элемент", maxElementForward(x, n))
		fmt.Println("Элемент обратный порядок", maxElementBackward(x, n))

		f := maxElementForward(x, n)
		b := maxElementBackward(x, n)

		if f > b {
			fmt.Println(f)
		} else {
			fmt.Println(b)
		}

	}

}
