package main

import "fmt"

// Узел бинарного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func main() {

	fmt.Println("Пример")
	arr := []int{-10, -3, 0, 5, 9}

	root := sortedArrayToBST(arr)
	fmt.Println("Вывод дерева:", root)
}

/*
Учитывая целочисленный массив nums, элементы которого отсортированы в порядке возрастания, преобразуйте
его в сбалансированный по высоте двоичное дерево поиска.

*/
