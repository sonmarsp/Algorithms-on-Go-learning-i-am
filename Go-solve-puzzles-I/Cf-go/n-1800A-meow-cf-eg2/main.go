package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Например, строки «meow», «mMmmEeOWwW», «MeOooOw» описывают мяуканье, а строки «Mweo»,
«MeO», «moew», «MmEW», «meowmeow» — нет.
1
8
MeoWMeoW
показывает да а должно нет запарывается на уникальности
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

		u := unique(wordP)
		fmt.Println("word")
		fmt.Println(u)
		fmt.Println("len")
		fmt.Println(len(u))

		var check string

		for _, v := range u {
			check += v
		}

		fmt.Println("check")
		fmt.Println(check)

		s := string("meow")
		fmt.Println(s)

		if len(u) == 4 {
			if s == check {
				fmt.Println("Yes")

			} else {
				fmt.Println("No")
			}
		} else {
			fmt.Println("No")
		}

	} //end main loop

}

func unique(intSlice []string) []string {

	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}

	}

	return list
}
