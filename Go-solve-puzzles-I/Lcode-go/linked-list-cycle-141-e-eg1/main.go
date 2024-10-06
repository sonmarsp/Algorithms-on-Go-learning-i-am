package main

import (
	"fmt"
)

//Для решения задачи определения наличия цикла в односвязном списке можно использовать
// алгоритм "черепаха и заяц" (также известный как алгоритм Флойда).
// Этот алгоритм использует два указателя: медленный (slow) и быстрый (fast).
//  Медленный указатель двигается на один узел за шаг, а быстрый — на два узла за шаг.
//  Если в списке есть цикл, то рано или поздно оба указателя встретятся.

// ListNode представляет узел односвязного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle проверяет, есть ли цикл в односвязном списке
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// Если медленный и быстрый указатели встретились, значит, есть цикл
		if slow == fast {
			return true
		}
	}

	// Если быстрый указатель достиг конца списка, значит, цикла нет
	return false
}

func main() {
	// Пример списка с циклом
	node4 := &ListNode{Val: -4}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 3, Next: node2}
	node4.Next = node2 // Создаем цикл, указывая на node2

	fmt.Println("Has cycle (with cycle):", hasCycle(node1)) // Ожидаемый вывод: true

	// Пример списка без цикла
	node6 := &ListNode{Val: 6}
	node5 := &ListNode{Val: 5, Next: node6}
	node4NoCycle := &ListNode{Val: 4, Next: node5}

	fmt.Println("Has cycle (no cycle):", hasCycle(node4NoCycle)) // Ожидаемый вывод: false

	// Пример пустого списка
	var emptyList *ListNode
	fmt.Println("Has cycle (empty list):", hasCycle(emptyList)) // Ожидаемый вывод: false
}
