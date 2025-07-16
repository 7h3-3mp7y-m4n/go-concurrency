package main

import (
	"fmt"
	"strconv"
	"time"
)

func bufferedExample() {
	fmt.Println("=== Buffered Channel Example ===")
	ch := make(chan string, 3) // Buffered channel

	for i := 1; i <= 3; i++ {
		msg := "Buffered " + strconv.Itoa(i)
		fmt.Println("Sending:", msg)
		ch <- msg // Doesn't block up to buffer size
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("Received:", <-ch)
	}
}

func unbufferedExample() {
	fmt.Println("\n=== Unbuffered Channel Example ===")
	ch := make(chan string) // Unbuffered channel

	go func() {
		for i := 1; i <= 3; i++ {
			msg := "Unbuffered " + strconv.Itoa(i)
			fmt.Println("Sending:", msg)
			ch <- msg // Blocks until received
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 1; i <= 3; i++ {
		fmt.Println("Received:", <-ch)
	}
}

func main() {
	bufferedExample()
	unbufferedExample()

	fmt.Println("\nConclusion:")
	fmt.Println("- Buffered channels are useful when you want to queue messages without blocking.")
	fmt.Println("- Unbuffered channels enforce synchronization between sender and receiver.")
}
