package main

import (
	"fmt"
	"os"
)

func main() {

	var w int

	//fmt.Print("Введите вес: ")
	fmt.Fscan(os.Stdin, &w)

	if w > 2 && w%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
