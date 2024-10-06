package main

import (
	"fmt"
)

// Используя математическую формулу без преобразования числа в строку
func main() {

	x := 121
	fmt.Println(isPalindrome(x)) // Output: true

	//x = -121
	//fmt.Println(isPalindrome(x)) // Output: false

	//x = 10
	//fmt.Println(isPalindrome(x)) // Output: false

	//x = 1221
	//fmt.Println(isPalindrome(x)) // Output: true

}

func isPalindrome(x int) bool {
	//Отрицательные числа и оканчивающиеся на ноль не могут быть палиндромами
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertNumber := 0
	//original := x

	for x > revertNumber {
		revertNumber = revertNumber*10 + x%10
		x /= 10
	}

	// Число является палиндромом, если оригинальная версия равна обратной версии
	// или если оригинальная версия равна половине обратной версии (в случае нечетного количества цифр)
	return x == revertNumber || x == revertNumber/10
}
