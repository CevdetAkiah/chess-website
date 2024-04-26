package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	vals := []int{1, 2, 3, 4, 5, 6}
	children := [][]int{
		{3, 2, 4},
		{5, 6},
		nil,
		nil,
		nil,
		nil,
	}

	root := buildTree(vals, children)
	fmt.Println(root)
}

// Function to build an n-ary tree
func buildTree(vals []int, children [][]int) *Node {
	nodes := make(map[int]*Node)

	// Create all nodes.
	for _, val := range vals {
		nodes[val] = &Node{Val: val}
	}

	// Link parent-child relationships.
	for i, childList := range children {
		parent := nodes[vals[i]]
		for _, childVal := range childList {
			child := nodes[childVal]
			parent.Children = append(parent.Children, child)
		}
	}

	return nodes[vals[0]]
}
