package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
	"unicode"
)

func main() {

	// Замеряем время выполнения
	start := time.Now()

	//Input: s = "A man, a plan, a canal: Panama"
	//Output: true
	//Explanation: "amanaplanacanalpanama" is a palindrome.

	//Убрать все знаки пунктуации и привести к нижнему регистру
	s := "A man, a plan, a canal: Panama"
	fmt.Println("Original string:", s)
	s = strings.ToLower(s)

	var result strings.Builder
	//Проходим по каждому символу в строке
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result.WriteRune(char)
		}

	}

	// Получаем итоговую строку из буфера
	cleanString := result.String()

	fmt.Println("Cleaned string:", cleanString)

	palindrome := true
	for i, j := 0, len(cleanString)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("i", i)
		fmt.Println("j", j)
		if cleanString[i] != cleanString[j] {
			palindrome = false
			break
		}

	}

	fmt.Println("Is palindrome:", palindrome)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
	// Получаем информацию об использовании памяти
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Выводим использование памяти
	fmt.Printf("Memory usage: Alloc = %v KB, TotalAlloc = %v KB, Sys = %v KB, NumGC = %v\n",
		memStats.Alloc/1024, memStats.TotalAlloc/1024, memStats.Sys/1024, memStats.NumGC)

}
