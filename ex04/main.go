package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func worker(id int, work <-chan int, res chan<- int) {
	for task := range work {
		fmt.Println("worker id: ", id, "   job: ", task)
		res <- task
	}
}

// в качестве способа завершения работы я решил выбрать контекст + канал
// контекст отлавливает сигналы ос для завершения работы
// после того как был получен сигнал в канал quit пишется значение
// как только в канал quit попадает какое то значение запуск горутин прекращается
func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	workers := 0
	fmt.Scan(&workers)
	jobs := workers
	work := make(chan int, workers)
	res := make(chan int, workers)

	quit := make(chan int)
	go func() {
		<-ctx.Done()
		fmt.Println("STOPPED BY PRINTING CTRL+C")
		quit <- 1
	}()
	for i := 0; i < workers; i++ {
		select {
		case <-quit:
			return
		default:
			go worker(i, work, res)
		}
	}
	for i := 0; i < jobs; i++ {
		work <- i
	}
	close(work)
	for i := 0; i < jobs; i++ {
		<-res
	}
}
