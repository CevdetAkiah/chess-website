package main

import (
	"fmt"
	"math/big"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	ix1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	ix2 := []int{5, 6, 4}

	l1 := AddNode(ix1)
	l2 := AddNode(ix2)

	AddLists(l1, l2)

}

func AddNode(ix []int) *ListNode {
	ll := []*ListNode{}
	for _, v := range ix {
		newNode := new(ListNode)
		newNode.Val = v
		ll = append(ll, newNode)
	}

	for i := 0; i < len(ll)-1; i++ {
		ll[i].Next = ll[i+1]
	}
	return ll[0]
}

func convertToSlice(node *ListNode) []int {
	var integers []int
	for iterator := node; iterator != nil; iterator = iterator.Next {
		integers = append(integers, iterator.Val)
	}
	return integers
}

// func sumSlice(xi []int) *big.Int {
// 	var sum = big.NewInt(0)
// 	var multiplier float64
// 	var addition uint64
// 	fmt.Println(len(xi))
// 	for i := 0; i < len(xi); i++ {
// 		multiplier = math.Pow(float64(10), float64(len(xi)-i-1))
// 		fmt.Printf("from the slice: %d\n", xi[i])
// 		fmt.Printf("multiplier %f\n", multiplier)

// 		if xi[i] == 0 {
// 			addition = uint64(multiplier)
// 			fmt.Printf("0 slice: %d", addition)
// 		} else {
// 			addition = uint64(xi[i]) * uint64(multiplier)
// 			sum += addition
// 			fmt.Printf("the sum: %d: %d\n", i, xi[i]*int(multiplier))
// 		}

// 	}

// 	return sum
// }

func sumSlice(xi []int) *big.Int {
	var sum = big.NewInt(0)
	multiplier := big.NewInt(1)
	for i := len(xi) - 1; i >= 0; i-- {
		xiBig := big.NewInt(int64(xi[i]))
		xiBig.Mul(xiBig, multiplier)
		sum.Add(sum, xiBig)
		multiplier.Mul(multiplier, big.NewInt(10))
	}
	return sum
}

// convert int to slice
func convertIntToSlice(sum *big.Int) []*big.Int {
	var xi []*big.Int
	divisor := big.NewInt(10)
	for sum.Cmp(big.NewInt(0)) != 0 {
		xi = append(xi, sum.Mod(sum, divisor))
		sum.Div(sum, divisor)
	}

	return xi
}

func AddLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := switchNodeOrder(l1)
	node2 := switchNodeOrder(l2)
	// for iterator := node1; iterator != nil; iterator = iterator.Next {
	// 	fmt.Println(iterator.Val)
	// }
	// for iterator := node2; iterator != nil; iterator = iterator.Next {
	// 	fmt.Println(iterator.Val)
	// }

	xi := convertToSlice(node1)
	xii := convertToSlice(node2)

	sum1 := sumSlice(xi)
	sum2 := sumSlice(xii)
	result := big.NewInt(0)
	result.Add(sum1, sum2)
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(result)
	bSlice := convertIntToSlice(result)
	fmt.Println(bSlice)
	// fmt.Println("result: ", xi)

	// break the sum back into nodes

	// reverse the node order
	return nil
}

func switchNodeOrder(l *ListNode) *ListNode {
	ll := []*ListNode{}
	for iterator := l; iterator != nil; iterator = iterator.Next {
		ll = append(ll, iterator)
	}
	// for _, v := range ll {
	// 	fmt.Println(v.Val)
	// }
	fmt.Println()
	ll2 := []*ListNode{}

	// switch the list order
	for i := len(ll) - 1; i >= 0; i-- {
		ll2 = append(ll2, ll[i])
	}
	// for _, v := range ll2 {
	// 	fmt.Println(v.Val)
	// }
	for i := 0; i < len(ll2); i++ {
		if i+1 == len(ll2) {
			ll2[i].Next = nil
			break
		}
		ll2[i].Next = ll2[i+1]
	}

	return ll2[0]
}

// func AddLists(l1 *ListNode, l2 *ListNode) *ListNode {

// 	if l1 == nil && l2 == nil {
// 		newNode := new(ListNode)
// 		return newNode
// 	}

// 	var check int
// 	for iterator := l1; iterator != nil; iterator = iterator.Next {
// 		check += iterator.Val
// 	}

// 	for iterator := l2; iterator != nil; iterator = iterator.Next {
// 		check += iterator.Val
// 	}

// 	if check == 0 {
// 		newNode := new(ListNode)
// 		fmt.Println("check: ", check)
// 		return newNode
// 	}
// 	var i1 []int
// 	var i2 []int

// 	for iterator := l1; iterator != nil; iterator = iterator.Next {
// 		i1 = append(i1, iterator.Val)
// 	}
// 	for iterator := l2; iterator != nil; iterator = iterator.Next {
// 		i2 = append(i2, iterator.Val)
// 	}

// 	// reverse slice order
// 	newslice1, newslice2 := func(ix1 []int, ix2 []int) ([]int, []int) {
// 		// var sum int
// 		var newslice1 []int
// 		var newslice2 []int
// 		for i := len(ix1) - 1; i > -1; i-- {
// 			newslice1 = append(newslice1, ix1[i])
// 		}

// 		for i := len(ix2) - 1; i > -1; i-- {
// 			newslice2 = append(newslice2, ix2[i])
// 		}

// 		return newslice1, newslice2
// 	}(i1, i2)

// 	// add slices to make one int
// 	sum := func(ix1 []int, ix2 []int) int {
// 		var sum int
// 		var multiplier float64
// 		for i := 0; i < len(ix1); i++ {
// 			multiplier = math.Pow(10, float64(len(ix1)-i-1))
// 			fmt.Println(ix1[i] * int(multiplier))
// 			sum += ix1[i] * int(multiplier)
// 		}
// 		for i := 0; i < len(ix2); i++ {
// 			multiplier = math.Pow(10, float64(len(ix2)-i-1))
// 			fmt.Println(ix2[i] * int(multiplier))

// 			sum += ix2[i] * int(multiplier)
// 		}
// 		return sum
// 	}(newslice1, newslice2)

// 	// convert int to slice
// 	finalSlice := func(ix1 int) []int {
// 		// var sum int
// 		var finalSlice []int
// 		for ix1 != 0 {
// 			finalSlice = append(finalSlice, ix1%10)
// 			ix1 /= 10
// 		}

// 		return finalSlice
// 	}(sum)

// 	var nodeList []*ListNode
// 	// convert each int to a node and append to nodeList
// 	for i := 0; i < len(finalSlice); i++ {
// 		num := finalSlice[i]
// 		newNode := new(ListNode)
// 		newNode.Val = num
// 		nodeList = append(nodeList, newNode)
// 	}

// 	for i := 0; i < len(nodeList)-1; i++ {
// 		if i+1 > len(nodeList) {
// 			continue
// 		} else {
// 			nodeList[i].Next = nodeList[i+1]
// 		}
// 	}

// 	return nodeList[0]
// }
