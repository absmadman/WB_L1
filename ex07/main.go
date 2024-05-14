package main

import (
	"fmt"
	"sync"
)

type syncMap struct {
	m map[int]int
	l sync.Mutex
}

func NewSM() *syncMap {
	return &syncMap{
		make(map[int]int),
		sync.Mutex{},
	}
}

func (s *syncMap) Add(key, val int) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[key] = val
}

func (s *syncMap) Remove(key int) {
	s.l.Lock()
	defer s.l.Unlock()
	delete(s.m, key)
}

func (s *syncMap) Get(key int) (int, bool) {
	s.l.Lock()
	defer s.l.Unlock()
	val, ok := s.m[key]
	return val, ok
}

// можно использовать syncMap но я напишу свою реализацию
func main() {
	sm := NewSM()
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()
			sm.Add(k, k)
		}()
	}
	wg.Wait()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()
			fmt.Println(sm.Get(k))
		}()
	}
	wg.Wait()
}
