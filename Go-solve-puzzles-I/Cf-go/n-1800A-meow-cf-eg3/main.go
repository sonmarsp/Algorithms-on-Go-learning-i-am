package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Например, строки «meow», «mMmmEeOWwW», «MeOooOw» описывают мяуканье, а строки «Mweo»,
«MeO», «moew», «MmEW», «meowmeow» — нет.
Полное решение но долго и много памяти 1466мс 4600кб
byte используют в решениях а у меня стринг может много памяти жрет
*/

func main() {

	var t int

	fmt.Fscan(os.Stdin, &t)
	fmt.Println(t)

	//Количество наборов
	for i := 0; i < t; i++ {

		//в принципе не нужна
		var n int
		fmt.Fscan(os.Stdin, &n)
		fmt.Println("n")
		fmt.Println(n)

		//получаем строку
		var w string
		fmt.Fscan(os.Stdin, &w)
		fmt.Println(w)

		var wordP []string
		for _, v := range w {

			//fmt.Println(k)
			p := string(v)

			lp := strings.ToLower(p)
			//fmt.Println(lp)
			wordP = append(wordP, lp)

		}
		fmt.Println(wordP)

		var word []string
		for i := 0; i < n-1; i++ {

			if wordP[i] != wordP[i+1] {

				word = append(word, wordP[i])
			}

		}
		word = append(word, wordP[n-1])
		fmt.Println(word)

		var check string

		for _, v := range word {
			check += v
		}

		fmt.Println("check")
		fmt.Println(check)

		s := string("meow")
		fmt.Println(s)

		if s == check {
			fmt.Println("Yes")

		} else {
			fmt.Println("No")
		}

	} //end main loop

}
