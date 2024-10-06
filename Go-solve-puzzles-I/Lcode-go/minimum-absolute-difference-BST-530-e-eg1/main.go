package main

import (
	"fmt"
	"math"
)

// Узел бинарного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция создания дерева из массива
func insertLevelOrder(arr []interface{}, i int) *TreeNode {
	var root *TreeNode
	if i < len(arr) && arr[i] != nil {
		val := arr[i].(int)
		root = &TreeNode{Val: val}
		root.Left = insertLevelOrder(arr, 2*i+1)
		root.Right = insertLevelOrder(arr, 2*i+2)
	}
	return root
}

// Функция создания дерева из массива
func createTreeFromArray(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	return insertLevelOrder(arr, 0)
}

// Функция для графического вывода дерева в консоль
func printTree(TreeNode *TreeNode, prefix string, isLeft bool) {
	if TreeNode != nil {
		fmt.Printf("%s", prefix)
		if isLeft {
			fmt.Printf("├── ")
			prefix += "│   "
		} else {
			fmt.Printf("└── ")
			prefix += "    "
		}
		fmt.Printf("%d\n", TreeNode.Val)
		printTree(TreeNode.Left, prefix, true)
		printTree(TreeNode.Right, prefix, false)
	}
}

func inOrderTraversal(TreeNode *TreeNode) {
	if TreeNode != nil {
		inOrderTraversal(TreeNode.Left)
		fmt.Print(TreeNode.Val, " ")
		inOrderTraversal(TreeNode.Right)
	}
}

func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	minDiff := math.MaxInt64

	var inOrder func(TreeNode *TreeNode)
	inOrder = func(TreeNode *TreeNode) {
		if TreeNode == nil {
			return
		}
		inOrder(TreeNode.Left)
		if prev != nil {
			diff := TreeNode.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = TreeNode
		inOrder(TreeNode.Right)
	}

	inOrder(root)
	return minDiff
}

func main() {

	fmt.Println("пример")

	arr := []interface{}{4, 2, 6, 1, 3}
	//arr := []interface{}{3, 9, 20, nil, 7, 17, 19}
	root := createTreeFromArray(arr)

	printTree(root, "", false)

	fmt.Println("In-order Traversal:")
	inOrderTraversal(root)
	fmt.Println()

	fmt.Println(getMinimumDifference(root))

}
