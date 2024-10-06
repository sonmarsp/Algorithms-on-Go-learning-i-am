package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
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

	re := regexp.MustCompile("[^a-z0-9]+")
	cleanString := re.ReplaceAllString(s, "")
	fmt.Println("Cleaned string:", cleanString)

	// Проверка на палиндром
	//isPalindrome := true
	//for i, j := 0, len(cleanedString)-1; i < j; i, j = i+1, j-1 {
	//	if cleanedString[i] != cleanedString[j] {
	//		isPalindrome = false
	//		break
	//	}
	//}

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
