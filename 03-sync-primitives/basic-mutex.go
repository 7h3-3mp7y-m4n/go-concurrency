package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	wg sync.WaitGroup
)

func main() {
	count := 0
	test()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // Lock the mutex before accessing the shared variable
			count++
			mu.Unlock() // Unlock the mutex to allow others to proceed
		}()
	}
	wg.Wait()
	fmt.Println("Final verdict ", count)
}

func test() {
	count := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final Count:", count)
}
