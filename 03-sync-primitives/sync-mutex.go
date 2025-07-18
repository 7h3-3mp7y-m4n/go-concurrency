package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	initFunc := func() {
		fmt.Println("Initialized only once.")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initFunc)
		}()
	}

	wg.Wait()
}
