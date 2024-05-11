package main

import (
	"fmt"
)

func typeAssertion(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case rune:
		return "rune"
	}
	return "unknown"
}

func main() {
	i1 := 0
	fmt.Println(typeAssertion(i1))
	i2 := "dasf"
	fmt.Println(typeAssertion(i2))
	i3 := 3.5
	fmt.Println(typeAssertion(i3))
	i4 := false
	fmt.Println(typeAssertion(i4))
	i5 := "😁"
	fmt.Println(typeAssertion(i5))
}
