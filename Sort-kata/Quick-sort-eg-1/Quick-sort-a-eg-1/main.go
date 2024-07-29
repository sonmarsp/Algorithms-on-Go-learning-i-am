package main

import (
	"fmt"
)

/*
*

  - Общая идея алгоритма состоит в следующем:
    Выбрать из массива элемент, называемый опорным. Это может быть любой из
    элементов массива. От выбора опорного элемента не зависит корректность
    алгоритма, но в отдельных случаях может сильно зависеть его эффективность

    Затем другие элементы в массиве перераспределяют так, чтобы элементы меньше
    опорного оказались до него, а большие или равные после. А дальше применяют
    рекурсивно первые два шага к подмассивам справа и слева от опорного массива.

    Недостатком считается unstable если числа повторяются две 8 8
*/
func main() {
	//start1 := time.Now()

	//arr := []int{3, 5, 6, 1, 7, 11, 14, 8, 18, 17, 43, 13}

	arr := []int{3, 5, 6, 1, 7, 9, 2, 4, 10, 8}

	fmt.Println(quickSortStart(arr))

}

func partition(arr []int, low, high int) ([]int, int) {
	//Отвечает за поиск опорной точки и перемещение всего в правильную сторону от
	// опорной точки
	pivot := arr[high]
	i := low
	fmt.Println(arr)
	//fmt.Println(pivot)
	//fmt.Println(low)
	//fmt.Println("Итерация")

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			//меняем местами

			//fmt.Println(pivot)
			//fmt.Println(low)
			//fmt.Println(arr[j])
			//fmt.Println("меняем местами")

			arr[i], arr[j] = arr[j], arr[i]
			i++
		}

	}

	fmt.Println(arr[i])
	fmt.Println(arr[high])

	arr[i], arr[high] = arr[high], arr[i]

	fmt.Println(arr[i])
	fmt.Println(arr[high])

	return arr, i

}

func quickSort(arr []int, low, high int) []int {

	if low < high {
		var p int

		//fmt.Println(low)
		//fmt.Println(high)
		//fmt.Println("Итерация")

		arr, p = partition(arr, low, high)

		//fmt.Println(arr)
		//fmt.Println(p)
		//fmt.Println("Итерация")
		//fmt.Println(arr)

		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
