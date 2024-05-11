package main

import (
	"fmt"
	"sync"
)

type SmartCounter struct {
	l     sync.Mutex
	count int
}

func NewSC() *SmartCounter {
	return &SmartCounter{
		count: 0,
	}
}

func (sc *SmartCounter) Increment() {
	sc.l.Lock()
	defer sc.l.Unlock()
	sc.count++
}

func (sc *SmartCounter) GetCount() int {
	return sc.count
}

func main() {
	sc := NewSC()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(sc.GetCount())
}
