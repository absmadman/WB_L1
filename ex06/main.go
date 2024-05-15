package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx, cancel := context.WithCancel(context.Background())
		wg.Add(1)
		go contx(ctx, wg)
		time.Sleep(2 * time.Second)
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Add(1)
		go ch(wg)
	}()

	wg.Wait()
}

// с помощью контекста
func contx(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("1st done")
			wg.Done()
			return
		}
	}
}

// с помощью канала
func ch(wg *sync.WaitGroup) {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("2nd done")
				wg.Done()
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
	done <- true
}
