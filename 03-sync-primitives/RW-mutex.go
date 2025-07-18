package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	var wg sync.WaitGroup
	data := 0

	wg.Add(1)
	//writer goroutine
	go func() {
		defer wg.Done()
		mu.Lock()
		data = 42
		mu.Unlock()
	}()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.RLock()
			fmt.Printf("Reader %d read value: %d\n", id, data)
			mu.RUnlock()
		}(i)
	}
	wg.Wait()
}
