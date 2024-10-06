package main

import "fmt"

// Узел бинарного дерева
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Функция для вставки узла в бинарное дерево поиска
func (n *Node) Insert(value int) {
	if value <= n.data {
		if n.left == nil {
			n.left = &Node{data: value}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: value}
		} else {
			n.right.Insert(value)
		}
	}
}

func maxDepth(root *Node) int {

	if root == nil {
		return 0
	}
	leftD := maxDepth(root.left)
	rightD := maxDepth(root.right)

	if leftD > rightD {
		return leftD + 1
	}

	return rightD + 1
}

func main() {

	fmt.Println("пример")

	// Пример 1
	root1 := &Node{data: 3}
	root1.Insert(9)
	root1.Insert(20)
	root1.left = nil // Это добавлено для соответствия примеру
	root1.right.left = &Node{data: 15}
	root1.right.right = &Node{data: 7}
	fmt.Println("Max Depth of Tree 1:", maxDepth(root1)) // Ожидаемый результат: 3

	// Пример 2
	root2 := &Node{data: 1}
	root2.right = &Node{data: 2}
	fmt.Println("Max Depth of Tree 2:", maxDepth(root2)) // Ожидаемый результат: 2
}
