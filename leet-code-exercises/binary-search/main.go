package main

import "fmt"

func main() {
	nums := []int{5}
	target := 5

	index := search(nums, target)

	fmt.Println(index)
}

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	if len(nums) == 1 {
		return 0
	}
	for low < high {
		median := (high + low) / 2
		if nums[median] == target {
			return median
		}
		if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return -1
}
