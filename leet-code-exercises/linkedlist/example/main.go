package main

import (
	"fmt"
	"strings"
)

// linked lists allow us to create a data structure where the element positions aren't determined by a continuous placement in the physical memory, as opposed to an array.

type linkedList struct {
	head *node
	len  int // length of ll
}

type node struct {
	value int
	next  *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func main() {
	ll := linkedList{}
	ll.add(1)
	ll.add(2)
	fmt.Println(ll)
	node := ll.get(1)
	fmt.Println("GOT NODE:", node.String())
	ll.remove(2)
	fmt.Println(ll)
}

func (ll *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value
	// newNode.next will point to nil as a pointer's nil value is nil

	// if ll is empty then make newNode the head
	if ll.head == nil {
		ll.head = newNode
	} else {
		// iterate through all the elements the head is connected to and grab the last element (the last element's next pointer will point to nil).
		iterator := ll.head
		for ; iterator.next != nil; iterator = iterator.next {
		}

		// set the last element's next to point to the newNode
		iterator.next = newNode
	}

	ll.len++
}

func (ll *linkedList) remove(value int) {
	// find node
	// set previous node next pointer to nil
	var previous *node
	for current := ll.head; current != nil; current = current.next {
		if current.value == value {
			if ll.head == current {
				ll.head = current.next
				return
			} else {
				previous.next = current.next
				return
			}
		}
		previous = current
	}

}
func (ll linkedList) get(value int) *node {
	for iter := ll.head; iter.next != nil; iter = iter.next {
		if iter.value == value {
			return iter
		}
	}
	return nil
}

func (ll linkedList) String() string {
	sb := strings.Builder{}
	for iter := ll.head; iter != nil; iter = iter.next {
		sb.WriteString(fmt.Sprintf("%s \n", iter.String()))
	}
	return sb.String()
}
