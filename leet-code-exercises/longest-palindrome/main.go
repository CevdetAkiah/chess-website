package main

import "fmt"

func main() {
	s := "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"
	fmt.Println("len: ", len(s))
	// palindrome := "naaan"
	// s := "a"
	p := longestPalindrome(s)
	fmt.Println(p)
}

func longestPalindrome(s string) int {
	runeMap := make(map[rune]int)
	length := 0
	oddCheck := 0

	for _, v := range s {
		runeMap[v]++
	}

	if len(runeMap) == 1 {
		return len(s)
	}

	for _, v := range runeMap {
		if v%2 == 0 {
			length += v
		} else {
			length += v - 1
			oddCheck++
		}
	}

	if oddCheck > 0 {
		length++
	}

	return length
}

// func longestPalindrome(s string) int {
// 	runeMap := make(map[rune]int)
// 	length := 0
// 	oddCheck := 0
// 	singleCharCheck := 0

// 	for _, v := range s {
// 		runeMap[v]++
// 	}
// 	if len(runeMap) == 1 {
// 		return len(s)
// 	}

// 	for _, v := range runeMap {
// 		if v == 1 {
// 			singleCharCheck++
// 		}

// 		if v%2 == 0 {
// 			length += v
// 		} else if v != 1 {
// 			length += v - 1
// 			oddCheck++
// 		}
// 	}

// 	if oddCheck == 1 && singleCharCheck == 0 {
// 		length++
// 	}
// 	if singleCharCheck != 0 {
// 		length++
// 	}

// 	return length
// }

// func longestPalindrome(s string) int {
// runeMap := make(map[rune]int)
// length := 0
// singleCharCheck := 0
// oddCheck := 0
// for _, v := range s {
// 	runeMap[v]++
// }
// if len(runeMap) == 1 {
// 	return runeMap[rune(s[0])]
// }

// if len(runeMap) == 2 {
// 	oddCheck := 0
// 	for _, v := range runeMap {
// 		if v%2 == 1 {
// 			oddCheck++
// 		}
// 	}
// 	if oddCheck == 2 {
// 		return len(s) - 1
// 	} else {
// 		return len(s)
// 	}

// }

// for _, v := range runeMap {
// 	if singleCharCheck == 0 {
// 		if v == 1 {
// 			singleCharCheck++
// 		}
// 	}

// 	if v%2 == 0 {
// 		length += v
// 	} else if v != 1 {
// 		length += v - 1
// 		oddCheck++
// 	}

// }

// if singleCharCheck != 0 {
// 	length++
// }
// if oddCheck == 1 && len {
// 	length++
// }

// return length
// }

// if len(runeMap) == 2 {
// 	for _, v := range runeMap {
// 		if v == 1 {
// 			length++
// 		}
// 		if v%2 == 0 {
// 			length += v
// 		} else if v%2 != 0 && v != 1 {
// 			length += (v - 1)
// 		}
// 	}
// 	return length
// }
