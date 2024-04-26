package main

import "fmt"

func main() {
	fmt.Println(recursive(10))
}

func recursive(i int) int {
	if i == 0 {
		return 0
	}

	return recursive(i - 1)
}
