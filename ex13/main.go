package main

import "fmt"

func main() {
	a := 5
	b := 10
	b, a = a, b
	fmt.Println(a, b)
}
