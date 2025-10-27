package main

import (
	"fmt"
	"sync"
)

type MutexCounter struct {
	mu    sync.Mutex
	count int
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	const numGoroutines = 100
	const incrementsPerGoroutine = 1000

	counter := MutexCounter{}
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	expected := numGoroutines * incrementsPerGoroutine
	fmt.Printf("Ожидаемое значение: %d, Итоговое значение: %d\n", expected, counter.Value())
}
