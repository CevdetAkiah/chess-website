package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	ix := []int{3, 2, 0, -4}
	ll := newList(ix)

	ll = setCycle(ll)

	ll2 := detectCycle(ll)

	fmt.Println(ll2.Val)

}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			n1, n2 := head, slow
			for n1 != n2 {
				n2 = n2.Next
				n1 = n1.Next
			}
			return n2
		}
	}
	return nil
}

// func detectCycle(head *ListNode) *ListNode {
// 	slow := head
// 	fast := head
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next

// 		if fast == slow {
// 			n1, n2 := head, slow
// 			for n1 != n2 {
// 				n2 = n2.Next
// 				n1 = n1.Next
// 			}
// 			return n2
// 		}
// 	}
// 	return nil
// }

// func detectCycle(head *ListNode) *ListNode {
// 	slow := head
// 	lMap := make(map[*ListNode]bool)

// 	for slow != nil && slow.Next != nil {
// 		if lMap[slow] {
// 			return slow
// 		}
// 		lMap[slow] = true
// 		slow = slow.Next

// 	}
// 	return nil
// }

func setCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast.Next = slow
	fmt.Println("cycle val: ", slow.Val)
	return head
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
