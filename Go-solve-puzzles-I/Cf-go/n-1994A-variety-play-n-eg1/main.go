package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение количества наборов входных данных
	tStr, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	for i := 0; i < t; i++ {
		// Чтение размеров матрицы
		nm, _ := reader.ReadString('\n')
		nmSplit := strings.Split(strings.TrimSpace(nm), " ")
		n, _ := strconv.Atoi(nmSplit[0])
		m, _ := strconv.Atoi(nmSplit[1])

		// Чтение матрицы a
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			rowStr, _ := reader.ReadString('\n')
			rowSplit := strings.Split(strings.TrimSpace(rowStr), " ")
			a[i] = make([]int, m)
			for j := 0; j < m; j++ {
				a[i][j], _ = strconv.Atoi(rowSplit[j])
			}
		}

		// Вывод матрицы a для проверки
		fmt.Println("Матрица a:")
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Printf("%d ", a[i][j])
			}
			fmt.Println()
		}
	}
}

/*
1
2 1
2
1

1
1 5
2 4 5 3 1

1
3 3
4 2 1
9 8 3
6 7 5
*/
