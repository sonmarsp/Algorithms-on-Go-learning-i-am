package main

import (
	"fmt"
)

func main() {

	//arr := []int{3, 5, 6, 1, 7, 11, 14, 8, 18, 17, 43, 34}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	//arr := []int{1, 2, 3}
	fmt.Println(arr)

	t := reverseArr(arr)
	fmt.Println(t)

}

func reverseArr(arr []int) []int {

	//s start arr
	//e end arr
	//получаем длину массива
	lp := len(arr)
	fmt.Println(lp)
	//число итераций
	var l int

	l = lp / 2
	fmt.Println("l")
	fmt.Println(l)

	e := lp - 1
	//fmt.Println(e)
	for s := 0; s < l; s++ {

		fmt.Println("s")
		fmt.Println(s)

		fmt.Println("e")
		fmt.Println(e)
		arr[s], arr[e] = arr[e], arr[s]

		//if s == e {
		//	break
		//}
		e--

	}

	return arr
}
