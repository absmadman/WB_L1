package main

import "fmt"

func ByTwoChan() {
	arr := []int{2, 4, 6, 8, 10}
	in := make(chan int)
	out := make(chan int)
	go func() {
		for _, num := range arr {
			in <- num
		}
		close(in)
	}()
	go func() {
		for num := range in {
			out <- (num * 2)
		}
		close(out)
	}()
	for num := range out {
		fmt.Print(num, " ")
	}
	fmt.Print("\n")
}

func main() {
	ByTwoChan()
}
