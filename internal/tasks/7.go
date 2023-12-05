package tasks

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m   map[string]int
	mux sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (s *SafeMap) Write(key string, value int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Read(key string) (int, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	value, ok := s.m[key]
	return value, ok
}

func Seventh() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		safeMap.Write("key1", 1)
	}()

	go func() {
		defer wg.Done()
		safeMap.Write("key2", 2)
	}()

	go func() {
		defer wg.Done()
		value, ok := safeMap.Read("key1")
		if ok {
			fmt.Println("Value for key1:", value)
		} else {
			fmt.Println("Key1 not found")
		}
	}()

	wg.Wait()
}
