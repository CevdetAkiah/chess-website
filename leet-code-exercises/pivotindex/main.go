package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	index := pivotIndex(nums)
	fmt.Println("FINAL INDEX: ", index)
}

// v1
// func pivotIndex(nums []int) int {
// 	var left []int
// 	var right []int

// 	for pi := 0; pi < len(nums); pi++ {
// 		for r := pi + 1; r < len(nums); r++ {
// 			right = append(right, nums[r])
// 		}
// 		if pi > 0 {
// 			for l := pi - 1; l < pi; l++ {
// 				left = append(left, nums[l])
// 			}
// 		}

// 		if sumSlice(left) == sumSlice(right) {
// 			return pi
// 		}
// 		right = []int{}
// 	}
// 	return -1
// }

// v2

func pivotIndex(nums []int) int {

	for pi := 0; pi < len(nums); pi++ {

		if sumSlice(nums, pi) {
			return pi
		}
	}
	return -1
}

func sumSlice(nums []int, index int) bool {
	var rightSum int
	var leftSum int

	for i := 0; i < index; i++ {
		leftSum += nums[i]
	}

	for i := index + 1; i < len(nums); i++ {
		rightSum += nums[i]
	}
	return leftSum == rightSum
}

// func pivotIndex(nums []int) int {
// 	fmt.Println(nums)
// 	leftSum := 0
// 	rightSum := 0
// 	for i := 0; i < len(nums); i++ {
// 		leftSum += nums[i]
// 		fmt.Println(nums[i])
// 		for j := i + 2; j < len(nums); j++ {
// 			rightSum += nums[j]
// 		}
// 		if rightSum == 0 {
// 			return 0
// 		}
// 		if leftSum == rightSum {
// 			i += 1
// 			return i
// 		}
// 		rightSum = 0
// 	}
// 	return -1
// }

// func pivotIndex(nums []int) int {
// 	fmt.Println(nums)
// 	leftSum := 0
// 	rightSum := 0
// 	for i := 0; i < len(nums); i++ {
// 		leftSum += nums[i]
// 		fmt.Println(nums[i])
// 		for j := len(nums) - 1; j > i+1; j-- {
// 			rightSum += nums[j]
// 		}
// 		if rightSum == 0 {
// 			return 0
// 		}
// 		if leftSum == rightSum {
// 			i += 1
// 			return i
// 		}
// 		rightSum = 0
// 	}
// 	return -1
// }
