package main

import (
	"fmt"
	"strings"
)

func reverseText(str string) string {
	split := strings.Split(str, " ")
	newStr := ""
	for i := 0; i < len(split); i++ {
		newStr = split[i] + " " + newStr
	}
	return newStr
}

func main() {
	str := "sun dog snow"
	fmt.Println(reverseText(str))
}
