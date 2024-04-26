package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSubsequence("", ""))
}

// func isSubsequence(s string, t string) bool {
// 	var check = make(map[rune]int)
// 	for i, v := range t {
// 		check[v] = i
// 	}

// 	for i, v := range s {
// 		if index, ok := check[v]; ok {
// 			if index < i {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

// func isSubsequence(s string, t string) bool {
// 	var check = make(map[rune]int)
// 	var ns string
// 	for i, v := range s {
// 		check[v] = i
// 	}
// 	for _, v := range t {
// 		if _, ok := check[v]; ok {
// 			ns += string(v)
// 		}
// 	}
// 	fmt.Println(ns)
// 	if strings.Contains(ns, s) {
// 		return true
// 	}
// 	return s == ns
// }

func isSubsequence(s string, t string) bool {
	var i int
	var j int
	if s == "" {
		return false
	}
	for {
		if i == len(s) {
			break
		}
		if j == len(t) {
			break
		}
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}

		if i == len(s) {
			return true
		}

	}
	return false
}
