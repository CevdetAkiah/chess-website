package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ix := []int{2, 4, 3}
	ix2 := []int{5, 6, 4}

	l1 := newList(ix)
	l2 := newList(ix2)

	l3 := addTwoNumbers(l1, l2)

	for iterator := l3; iterator != nil; iterator = iterator.Next {
		fmt.Println(iterator.Val)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var prevNode *ListNode
	var nextNode *ListNode

	currentNode := l1
	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	return prevNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var head *ListNode
// 	var tail *ListNode

// 	for l1 != nil {
// 		newNode := &ListNode{Val: l1.Val + l2.Val}

// 		if head == nil {
// 			head = newNode
// 			tail = newNode
// 		} else {
// 			tail.Next = newNode
// 			tail = newNode
// 		}
// 		l1 = l1.Next
// 		l2 = l2.Next

// 	}

// 	return head
// }

func newList(ix []int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for _, v := range ix {
		newNode := &ListNode{Val: v}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}

	return head
}
