package main

import (
	"fmt"
	"sync"
	"time"
)

func memoryBenchmark() {
	now := time.Now()

	var i = 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for range 10000000 {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			i++
		}()
	}

	wg.Wait()

	diff := time.Since(now)

	fmt.Println(diff.Milliseconds())
}
