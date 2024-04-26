package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ix1 := []int{1, 2, 3, 4}

	list := newList(ix1)
	rList := reverseList(list)

	for ; rList != nil; rList = rList.Next {
		fmt.Println(rList.Val)
	}
}

func reverseList(head *ListNode) *ListNode {
	var prevNode *ListNode
	current := head
	for current != nil {
		nextNode := current.Next
		current.Next = prevNode
		prevNode = current
		current = nextNode
	}

	return prevNode
}

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
