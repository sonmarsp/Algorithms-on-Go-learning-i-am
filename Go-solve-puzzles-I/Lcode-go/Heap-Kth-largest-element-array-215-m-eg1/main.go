package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// swap обмен элементов местами в куче
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	// метод библиотеки heap.Init  container/heap
	heap.Init(h)

	// Добовляем элементы в кучу
	for _, num := range nums {
		heap.Push(h, num)
		// Если размер кучи превышает k, удаляем минимальный элемент
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

func main() {

	fmt.Println("пример")
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	fmt.Println(findKthLargest(nums1, k1)) // Вывод: 5

	//nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//k2 := 4
	//fmt.Println(findKthLargest(nums2, k2)) // Вывод: 4
}

/*
Вход: nums = [3, 2, 1, 5, 6, 4], k = 2
После сортировки: [6, 5, 4, 3, 2, 1]
k-й по величине элемент: 5 (второй элемент после сортировки)
Выход: 5
*/
