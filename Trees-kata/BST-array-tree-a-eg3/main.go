package main

import (
	"fmt"
)

// Узел бинарного дерева
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Функция для создания дерева из массива
func createTreeFromArray(arr []interface{}) *Node {
	if len(arr) == 0 {
		return nil
	}
	return insertLevelOrder(arr, 0)
}

// Правый ребенок текущего узла определяется путем вычисления 2*i+2.
// Эта формула соответствует схеме обхода по порядку уровней,
// где левый дочерний узел находится по индексу 2*i+1,
// а правый - по индексу 2*i+2.
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

func learnTraversals(node *Node) {
	if node != nil {
		fmt.Println(node.data, " ")
		learnTraversals(node.left)
		learnTraversals(node.right)
	}

}

func main() {
	fmt.Println("дерево")

	arr := []interface{}{3, 9, 20, nil, nil, 15, 7}
	//arr := []interface{}{3, 9, 20, 8, 7, 17, 19}
	root := createTreeFromArray(arr)

	fmt.Println("Как мне увидеть узлы связь этих nodes?")
	printTree(root, "", false)

	//fmt.Println(root.left)
	//fmt.Println(root.left.right)

	learnTraversals(root)

}
