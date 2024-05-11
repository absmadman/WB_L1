package main

import (
	"fmt"
	"strings"
)

// не может работать с рунами но занимает меньше памяти
func checkByArray(str string) bool {
	str = strings.ToLower(str)
	var abc [26]int
	for _, c := range str {
		abc[c-'a']++
		if abc[c-'a'] > 1 {
			return false
		}
	}
	return true
}

// может работать с рунами но занимает много памяти
func checkByMap(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]int, 0)
	for _, c := range str {
		m[c]++
		if m[c] > 1 {
			return false
		}
	}
	return true
}

func main() {
	a := "abcd"
	b := "abCdefAaf"
	c := "aabcd"
	fmt.Println(checkByArray(a))
	fmt.Println(checkByArray(b))
	fmt.Println(checkByArray(c))
	fmt.Println(checkByMap(a))
	fmt.Println(checkByMap(b))
	fmt.Println(checkByMap(c))
}
