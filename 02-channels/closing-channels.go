package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close channel when done sending
	}()
	for value := range ch {
		fmt.Println("Received <===", value)
	}
	fmt.Println("All values received. Channel is closed.")
}
