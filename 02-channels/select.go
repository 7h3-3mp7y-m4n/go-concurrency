package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	unbuffCh := make(chan string)
	bufferedCh := make(chan string, 1)
	bufferedCh <- "Buffered value" // This sends immediately
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		unbuffCh <- "Unbuffered value"
	}()
	select {
	case msg := <-bufferedCh:
		fmt.Println("Received from buffered channel ====>", msg)
	case msg := <-unbuffCh:
		fmt.Println("Received from unbuffered channel =====>", msg)
	}
}
