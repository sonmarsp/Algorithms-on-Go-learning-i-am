package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := atoi(scanner.Text())

	results := make([]string, t)

	for i := 0; i < t; i++ {
		scanner.Scan() // Считываем n, но не используем его
		scanner.Scan()
		s := scanner.Text()

		if canTransformToOne(s) {
			results[i] = "Yes"
		} else {
			results[i] = "No"
		}
	}

	fmt.Println(strings.Join(results, "\n"))
}

func canTransformToOne(s string) bool {
	count0, count1 := 0, 0
	for _, char := range s {
		if char == '0' {
			count0++
		} else {
			count1++
		}
	}
	// Если количество единиц больше количества нулей, можно преобразовать в [1]
	return count1 > count0
}

func atoi(s string) int {
	result := 0
	for _, char := range s {
		result = result*10 + int(char-'0')
	}
	return result
}

//неверное решение, непонятное условие задачи
