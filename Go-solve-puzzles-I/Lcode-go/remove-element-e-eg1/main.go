package main

import (
	"fmt"
)

func main() {

	nums := []int{3, 2, 2, 3}
	needle := 3

	//fmt.Println("testo")
	//fmt.Println(nums)

	//fmt.Println(needle)
	//fmt.Println(nums[:2])

	c, res := remove(nums, needle)

	//stats := new(runtime.MemStats)
	//runtime.ReadMemStats(stats)
	//fmt.Printf("stats: %d\n", stats.HeapAlloc)

	res2 := reverseArr(res)
	res2 = res2[:len(res2)-c]

	fmt.Println(c)
	fmt.Println(res)
	fmt.Println(res2[2:4])

	fmt.Println("nums--")
	fmt.Println(nums)

}

func remove(sl []int, n int) (int, []int) {

	var nsl []int

	for _, v := range sl {
		if v != n {
			nsl = append(nsl, v)
		}
	}

	sl = append(sl, nsl...)

	//fmt.Println("nsl")
	//fmt.Println(nsl)
	//fmt.Println(sl)

	return len(nsl), sl
}

func reverseArr(arr []int) []int {

	dim := len(arr)

	for i := range arr[:dim/2] {

		arr[i], arr[dim-i-1] = arr[dim-i-1], arr[i]
	}

	return arr

}
