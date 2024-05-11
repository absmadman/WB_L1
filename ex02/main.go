package main

import (
	"fmt"
)

// С помощью одного канала
func ByChan() {
	arr := []int{2, 4, 6, 8, 10}
	res := make(chan int)
	for i := range arr {
		go func() {
			res <- (arr[i] * arr[i])
		}()
		fmt.Print(<-res, " ")
	}
	fmt.Print("\n")
}

// С помощью двух каналов
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
			out <- (num * num)
		}
		close(out)
	}()
	for num := range out {
		fmt.Print(num, " ")
	}
	fmt.Print("\n")
}

func main() {
	ByChan()
	ByTwoChan()
}
