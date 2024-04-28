package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ix1 := []int{1, 2, 4}
	ix2 := []int{1, 3, 4}

	node1 := newList(ix1)
	node2 := newList(ix2)

	node3 := mergeTwoLists(node1, node2)

	for iterator := node3; iterator != nil; iterator = iterator.Next {
		fmt.Println(iterator.Val)
	}
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
			tail.Next = newNode // previous tail.Next points to the new node
			tail = newNode      // new node becomes the new tail of the linked list
		}
	}
	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	for list1 != nil && list2 != nil {
		var node *ListNode

		if list1.Val < list2.Val {
			node = list1
			list1 = list1.Next
		} else {
			node = list2
			list2 = list2.Next
		}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}

	}

	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}

	return head
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	var head *ListNode
// 	var tail *ListNode
// 	for list1 != nil && list2 != nil {
// 		var node *ListNode

// 		if head == nil {
// 			if list1.Val < list2.Val {
// 				head = list1
// 				tail = list1
// 			} else {
// 				head = list2
// 				tail = list2
// 			}
// 		} else {
// 			if tail == list1 {
// 				tail.Next = list2
// 				tail = list2
// 				list2 = list2.Next
// 			} else {
// 				tail.Next = list1
// 				tail = list1
// 				list1 = list1.Next
// 			}
// 		}

// 	}

// 	return head
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	var head *ListNode
// 	var tail *ListNode

// 	for list1 != nil && list2 != nil {
// 		var node *ListNode
// 		if list1.Val < list2.Val {
// 			node = list1
// 			list1 = list1.Next
// 		} else {
// 			node = list2
// 			list2 = list2.Next
// 		}

// 		if head == nil {
// 			head = node
// 			tail = node
// 		} else {
// 			tail.Next = node
// 			tail = node
// 		}
// 	}

// 	// Append any remaining nodes from list1 or list2
// 	if list1 != nil {
// 		if head == nil {
// 			return list1
// 		}
// 		tail.Next = list1
// 	} else if list2 != nil {
// 		if head == nil {
// 			return list2
// 		}
// 		tail.Next = list2
// 	}

// 	return head
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	ll := []*ListNode{}
// 	newNode := new(ListNode)
// 	iterator1 := list1
// 	iterator2 := list2
// 	for {
// 		newNode = iterator1
// 		newNode.Next = iterator2
// 		if iterator2.Next == nil {
// 			break
// 		}
// 		ll = append(ll, newNode)

// 		iterator1 = iterator1.Next
// 		iterator2 = iterator2.Next
// 	}
// 	return ll[0]
// }

// ix1 := []int{1, 2, 4}
// ix2 := []int{1, 3, 4}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	ll := []*ListNode{}
// 	newNode := new(ListNode)
// 	iterator2 := list2
// 	for iterator := list1; iterator != nil && iterator2 != nil; iterator = iterator.Next {
// 		newNode = iterator
// 		iterator2.Next = iterator.Next
// 		newNode.Next = iterator2
// 		ll = append(ll, newNode)
// 		iterator2 = iterator2.Next
// 	}
// 	return ll[0]
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	ll := []int{}
// 	iterator2 := list2
// 	for iterator := list1; iterator != nil && iterator2 != nil; iterator = iterator.Next {
// 		ll = append(ll, iterator.Val, iterator2.Val)
// 		iterator2 = iterator2.Next
// 	}
// 	newNode := newList(ll)

// 	return newNode
// }
