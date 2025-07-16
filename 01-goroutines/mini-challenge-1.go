package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		// Capture the loop variable
		go func(workerID int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", workerID)
		}(i)
	}

	wg.Wait()
}
