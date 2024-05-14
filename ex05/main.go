package main

import (
	"fmt"
	"time"
)

const n = 2

func main() {
	out := make(chan int)
	start := time.Now()
	timer := time.NewTimer(time.Second * time.Duration(n))
	go func() {
		for i := 0; ; i++ {
			out <- i
			time.Sleep(time.Millisecond * 25)
		}
	}()
	for num := range out {
		select {
		case <-timer.C:
			now := time.Now()
			fmt.Printf("STOPPED AFTER %s", now.Sub(start))
			return
		default:
			fmt.Println("HARD WORKING WITH NUMBER: ", num)
		}
	}
}
