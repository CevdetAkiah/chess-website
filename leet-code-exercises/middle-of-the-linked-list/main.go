package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	i := []int{12, 86, 47, 6, 23, 6, 11, 30, 16, 81, 62, 32, 80, 61, 66, 41, 8, 88, 5, 98, 77, 54, 24, 60, 52, 32, 99, 84, 81, 66, 1, 25, 31, 27, 70, 90, 19, 54, 50, 6, 72, 32, 69, 88, 18, 10, 75, 40, 22, 97}
	ix := newList(i)

	mid := middleNode(ix)
	fmt.Println(mid.Val)
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// func middleNode(head *ListNode) *ListNode {
// 	var midNode *ListNode
// 	var count int
// 	var middle int
// 	iterator := head
// 	for ; iterator != nil; iterator = iterator.Next {
// 		count++
// 	}
// 	fmt.Println("count: ", count)
// 	middle = count/2 + 1

// 	fmt.Println("Middle: ", middle)
// 	count = 0
// 	iterator = head
// 	for ; iterator != nil; iterator = iterator.Next {
// 		count++
// 		if count == middle {
// 			midNode = iterator
// 			return midNode
// 		}
// 	}

// 	return nil
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

// func middleNode(head *ListNode) *ListNode {
// 	var midNode *ListNode
// 	var count int
// 	var middle int
// 	iterator := head
// 	for ; iterator != nil; iterator = iterator.Next {
// 		count++
// 	}
// 	fmt.Println("count: ", count)
// 	middle = count / 2

// 	if middle%2 == 0 || middle == 3 || middle == 5 || middle == 7 {
// 		middle = middle + 1
// 	}
// 	if count == 2 || count == 3 {
// 		middle = 2
// 	}
// 	fmt.Println("Middle: ", middle)
// 	count = 0
// 	iterator = head
// 	for ; iterator != nil; iterator = iterator.Next {
// 		count++
// 		if count == middle {
// 			midNode = iterator
// 			return midNode
// 		}
// 	}

// 	return nil
// }

// func middleNode(head *ListNode) *ListNode {
// 	slow := head
// 	fast := head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}

// 	return slow
// }
