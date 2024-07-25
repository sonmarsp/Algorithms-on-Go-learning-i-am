package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}

	//arr := []int{1, 2, 3}

	fmt.Println(arr)
	res := make([]int, len(arr))
	fmt.Println(res)
	copy(res[3:], arr[:len(arr)-3])
	fmt.Println(res)
	copy(res[:], arr[len(arr)-3:])
	fmt.Println(res)

}
