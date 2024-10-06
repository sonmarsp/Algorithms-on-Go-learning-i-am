package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Чтение количества наборов входных данных
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		// Чтение размеров матрицы
		scanner.Scan()
		nm := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(nm[0])
		m, _ := strconv.Atoi(nm[1])

		// Чтение матрицы a
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			row := strings.Split(scanner.Text(), " ")
			a[i] = make([]int, m)
			for j := 0; j < m; j++ {
				a[i][j], _ = strconv.Atoi(row[j])
			}
		}

		// Вывод матрицы a для проверки
		/*fmt.Println("Матрица a:")
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Printf("%d ", a[i][j])
			}
			fmt.Println()
		}
		*/

		// Проверка на невозможность
		if n == 1 && m == 1 {
			fmt.Println(-1)
			continue
		}

		// Создание одномерного массива чисел от 1 до n * m
		nums := make([]int, n*m)
		for i := 0; i < n*m; i++ {
			nums[i] = i + 1
		}

		// Перемешивание чисел так, чтобы они не совпадали с матрицей a
		b := make([][]int, n)
		used := make([]bool, n*m+1)
		for i := 0; i < n; i++ {
			b[i] = make([]int, m)
			for j := 0; j < m; j++ {
				for _, num := range nums {
					if num != a[i][j] && !used[num] {
						b[i][j] = num
						used[num] = true
						break
					}
				}
			}
		}

		// Проверка на невозможность
		possible := true
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if a[i][j] == b[i][j] {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}

		if !possible {
			fmt.Fprintln(writer, -1)
		} else {
			// Вывод матрицы b
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					fmt.Fprintf(writer, "%d ", b[i][j])
				}
				fmt.Fprintln(writer)
			}
		}
	}
}

/*
1
1 1
1


1
2 1
2
1

1
1 5
2 4 5 3 1

1
2 4
1 2 3 4
5 6 7 8

1
3 3
4 2 1
9 8 3
6 7 5
*/
