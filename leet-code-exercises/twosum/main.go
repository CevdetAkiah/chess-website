package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var tt []int

	for i := 0; i < len(nums); i++ {
		if len(tt) == 2 {
			break
		}
		// if math.Abs(float64(nums[i])) > math.Abs(float64(target)) {
		// 	i++
		// 	if i > len(nums)-1 {
		// 		break
		// 	}
		// }
		test := nums[i]
		for j := 0; j < len(nums); j++ {
			if i == j {
				j++
				if j > len(nums)-1 {
					break
				}
			}
			test2 := nums[j]
			if test+test2 == target {
				tt = append(tt, i, j)
				break
			}

		}
	}

	return tt

}

func main() {
	nums := []int{-3, 4, 3, 90}
	target := 0
	result := twoSum(nums, target)
	fmt.Println(result)
}
