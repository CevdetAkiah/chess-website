package main

import "fmt"

func main() {

	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}

func runningSum(nums []int) []int {
	sumList := []int{}
	sumList = append(sumList, nums[0])
	for i := 1; i < len(nums); i++ {
		sum := sumList[i-1] + nums[i]
		fmt.Println(sum)
		sumList = append(sumList, sum)
	}
	return sumList
}
