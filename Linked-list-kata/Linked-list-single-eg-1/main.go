package main

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddToFront(data string) {

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) String() string {
	result := ""
	current := list.head
	for current != nil {
		result += current.data + " -> "
		current = current.next
	}
	result += "nil"
	return result
}

func (list *LinkedList) AddToEnd(data string) {

	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode

}

func main() {

	fmt.Println("пример")
	list := &LinkedList{}

	list.AddToFront("first")
	list.AddToFront("second")
	list.AddToEnd("end")

	fmt.Println(list)

}
