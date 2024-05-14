package main

import "fmt"

// С помощью одного канала
func ByChan() {
	arr := []int{2, 4, 6, 8, 10}
	res := make(chan int)
	sum := 0
	for i := range arr {
		go func() {
			res <- (arr[i] * arr[i])
		}()
		sum += <-res
	}
	fmt.Println(sum)
}

// С помощью двух каналов
func ByTwoChan() {
	arr := []int{2, 4, 6, 8, 10}
	in := make(chan int)
	out := make(chan int)
	sum := 0
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
		sum += num
	}
	fmt.Println(sum)
}

func main() {
	ByChan()
	ByTwoChan()
}
