package main

import (
	"fmt"
)

type CircularNode struct {
	data string
	next *CircularNode
}

type CircularLinkedList struct {
	head *CircularNode
}

// Пример добавления элемента в конец кольцевого односвязного списка
func (list *CircularLinkedList) AddToEnd(data string) {
	newNode := &CircularNode{data: data}
	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
	} else {
		temp := list.head
		for temp.next != list.head {
			temp = temp.next
		}
		temp.next = newNode
		newNode.next = list.head
	}
}

func (list *CircularLinkedList) String() string {
	if list.head == nil {
		return "nil"
	}
	result := ""
	current := list.head
	for {
		result += fmt.Sprintf("%s -> ", current.data)
		current = current.next
		if current == list.head {
			break
		}
	}
	result += fmt.Sprintf("%s (head)", list.head.data)
	return result
}

func main() {

	fmt.Println("пример")
	list := &CircularLinkedList{}

	list.AddToEnd("first")
	list.AddToEnd("second")
	list.AddToEnd("third")

	fmt.Println(list)

}
