package main

import (
	"fmt"
	"time"
)

func main() {
	go Hello() // here we start our new goroutine
	// we will give some time to it otherwise main will exit without getting any output from our go routine
	go HelloCount()
	time.Sleep(time.Second * 1)
	fmt.Println("Main function ends!")
}

func Hello() {
	fmt.Println("hello from goroutine!")
}

// we can see that after printing hello x times we exit as our main code because we add time.sleep to buy some time for main
func HelloCount() {
	for i := 1; i < 10; i++ {
		fmt.Printf("Hello %d times \n", i)
		if i == 3 {
			time.Sleep(time.Second * 2)
		}
	}
}
