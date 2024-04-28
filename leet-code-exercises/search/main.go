package main

import (
	"fmt"
)

func main() {

	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(9, items))

}

// func BenchmarkLinearSearch(b *testing.B) {
// 	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}

// 	for i := 0; i < b.N; i++ {
// 		linearSearch(items, 58)
// 	}

// }

// func linearSearch(dataList []int, key int) bool {
// 	for _, index := range dataList {
// 		if key == index {
// 			return true
// 		}
// 	}

// 	return false
// }

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

// func BenchmarkBinarySearch(b *testing.B) {
// 	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}

// 	for i := 0; i < b.N; i++ {
// 		binarySearch(58, items)
// 	}

// }
