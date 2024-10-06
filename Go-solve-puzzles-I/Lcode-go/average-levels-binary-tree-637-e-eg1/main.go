package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

// Функция для графического вывода дерева в консоль
func printTree(node *Node, prefix string, isLeft bool) {
	if node != nil {
		fmt.Printf("%s", prefix)
		if isLeft {
			fmt.Printf("├── ")
			prefix += "│   "
		} else {
			fmt.Printf("└── ")
			prefix += "    "
		}
		fmt.Printf("%d\n", node.data)
		printTree(node.left, prefix, true)
		printTree(node.right, prefix, false)
	}
}

func averageOfLevels(root *Node) []float64 {
	if root == nil {
		return []float64{}
	}

	result := []float64{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		sum := 0
		count := queue.Len()

		for i := 0; i < count; i++ {
			element := queue.Front()
			node := element.Value.(*Node)
			queue.Remove(queue.Front())

			sum += node.data

			if node.left != nil {
				queue.PushBack(node.left)
			}

			if node.right != nil {
				queue.PushBack(node.right)
			}

		}

		average := float64(sum) / float64(count)
		result = append(result, average)
	}

	return result
}

// Функция для создания дерева из массива
func createTreeFromArray(arr []interface{}) *Node {
	if len(arr) == 0 {
		return nil
	}
	return insertLevelOrder(arr, 0)
}

func insertLevelOrder(arr []interface{}, i int) *Node {
	var root *Node
	if i < len(arr) && arr[i] != nil {
		val := arr[i].(int)
		root = &Node{data: val}
		root.left = insertLevelOrder(arr, 2*i+1)
		root.right = insertLevelOrder(arr, 2*i+2)
	}
	return root
}

func main() {
	arr := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := createTreeFromArray(arr)

	printTree(root, "", false)

	fmt.Println("Average of levels:", averageOfLevels(root))

}
