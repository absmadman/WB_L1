package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(str string) string {
	newStr := ""
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		newStr = string(r) + newStr
		i += size
	}
	return newStr
}

func main() {
	str := "㆒㆓㆔"
	fmt.Println(reverseString(str))
}
