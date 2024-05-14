package main

import "fmt"

// своеобразный multiset
func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]int)
	for _, str := range strs {
		set[str]++
	}
	fmt.Println(set)
}
