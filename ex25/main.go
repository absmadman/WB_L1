package main

import (
	"log"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	log.Println("START")
	sleep(3 * time.Second)
	log.Println("END")
}
