package main

import "fmt"

func main() {
	s, t := "egg", "add"
	ss, tt := "foo", "bar"
	sss, ttt := "badc", "baba"
	ssss, tttt := "paper", "title"

	fmt.Println(isIsomorphic(s, t))
	fmt.Println(isIsomorphic(ss, tt))
	fmt.Println(isIsomorphic(sss, ttt))
	fmt.Println(isIsomorphic(ssss, tttt))

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sMap = make(map[byte]byte)
	var tMap = make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if val, ok := sMap[s[i]]; ok {
			if val != t[i] {
				return false
			}
		}
		if val, ok := tMap[t[i]]; ok {
			if val != s[i] {
				return false
			}
		}
		sMap[s[i]] = t[i]
		tMap[t[i]] = s[i]
	}
	return true
}
