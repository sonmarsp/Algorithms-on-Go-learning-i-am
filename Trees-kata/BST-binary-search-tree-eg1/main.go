package main

import (
	"container/list"
	"fmt"
)

// Узел бинарного дерева
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Вставка узла в бинарное дерево
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

// Прямой обход (Pre-order Traversal)
// Обходит дерево в порядке "корень -> левое поддерево -> правое поддерево
func preOrderTraversal(node *Node) {
	if node != nil {
		fmt.Print(node.data, " ")
		preOrderTraversal(node.left)
		preOrderTraversal(node.right)
	}
}

// Центрированный обход (In-order Traversal)
// Обходит дерево в порядке "левое поддерево -> корень -> правое поддерево
func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.left)
		fmt.Print(node.data, " ")
		inOrderTraversal(node.right)
	}
}

// Post-order Traversal: Обходит дерево в порядке "левое поддерево -> правое поддерево -> корень"
func postOrderTraversal(node *Node) {
	if node != nil {
		postOrderTraversal(node.left)
		postOrderTraversal(node.right)
		fmt.Print(node.data, " ")
	}
}

// Level-order Traversal: Обходит дерево по уровням слева направо, используя очередь для хранения узлов текущего уровня.
func levelOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	queue := list.New()
	queue.PushBack(node)

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(*Node)
		fmt.Print(node.data, " ")

		if node.left != nil {
			queue.PushBack(node.left)
		}

		if node.right != nil {
			queue.PushBack(node.right)
		}

		queue.Remove(element)
	}

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

func main() {
	fmt.Println("дерево")

	root := &Node{data: 60}
	root.Insert(40)
	root.Insert(20)
	root.Insert(10)
	root.Insert(30)
	root.Insert(50)
	root.Insert(80)
	root.Insert(70)
	root.Insert(90)

	fmt.Println("Root Node:", root)
	fmt.Println("Left Child of Root:", root.left)
	fmt.Println("Right Child of Root:", root.right)
	fmt.Println("Left Child of Left Child:", root.left.left)

	fmt.Println("Pre-order Traversal:")
	preOrderTraversal(root)
	fmt.Println()

	fmt.Println("In-order Traversal:")
	inOrderTraversal(root)
	fmt.Println()

	fmt.Println("Post-order Traversal:")
	postOrderTraversal(root)
	fmt.Println()

	fmt.Println("Level-order Traversal:")
	levelOrderTraversal(root)
	fmt.Println()

	fmt.Println("Graphical Representation of the Tree:")
	printTree(root, "", false)

}
