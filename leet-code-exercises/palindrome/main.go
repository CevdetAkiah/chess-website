package main

import "fmt"

func main() {
	toBreak := 411114
	fmt.Println(palindromChecker(toBreak))
}

func palindromChecker(toBreak int) bool {
	if toBreak < 0 {
		return false
	}
	toCheck := breakInt(toBreak)
	j := len(toCheck) - 1
	for i := 0; i < len(toCheck)-1; i++ {
		if j <= (len(toCheck)-1)/2 {
			break
		}
		if toCheck[i] != toCheck[j] {
			return false
		}
		j--

	}
	return true
}

func breakInt(toBreak int) []int {
	var slice []int
	for toBreak != 0 {
		slice = append(slice, toBreak%10)
		toBreak /= 10
	}
	fmt.Println(len(slice))

	return slice
}
