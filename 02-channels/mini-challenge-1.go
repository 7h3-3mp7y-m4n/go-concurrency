package main

import (
	"fmt"
	"strconv"
)

func main() {
	messageChannel := make(chan string)
	go func() {
		for i := 1; i <= 5; i++ {
			messageChannel <- "hello " + strconv.Itoa(i)
		}
		close(messageChannel)
	}()
	for value := range messageChannel {
		fmt.Println("Received <=== ", value)
	}
	fmt.Println("DONE!!")
}
