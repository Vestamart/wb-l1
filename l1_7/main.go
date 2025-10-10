package main

import (
	"fmt"
	"sync"
)

func main() {
	var concurrentMap sync.Map
	var wg sync.WaitGroup
	numGoroutines := 100
	key := "hits"

	concurrentMap.Store(key, 0)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				if val, ok := concurrentMap.Load(key); ok {
					current := val.(int)
					concurrentMap.Store(key, current+1)
				}
			}
		}()
	}

	wg.Wait()

	finalVal, _ := concurrentMap.Load(key)
	expectedTotal := numGoroutines * 1000

	fmt.Printf("Expected total value: %d\n", expectedTotal)
	fmt.Printf("Actual total value:%d\n", finalVal.(int))
}
