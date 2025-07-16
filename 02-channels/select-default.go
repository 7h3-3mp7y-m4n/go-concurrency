package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan string)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		ch <- "Message arrived"
	}()
	select {
	case msg := <-ch:
		fmt.Println("Received ===>", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No message within 2 seconds")
	default:
		fmt.Println("Nothing ready immediately")
	}
}
