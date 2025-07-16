package main

import (
	"fmt"
	"sync"
	"time"
)

// ading wait group
func main() {
	var wg sync.WaitGroup
	wg.Add(2) //expecting 2 goroutine
	go func() {
		defer wg.Done()
		waitHello()
	}()
	go func() {
		defer wg.Done()
		helloWaitGroup()
	}()
	wg.Wait()
	fmt.Println("we waited for main and for goroutine!!")
}

func helloWaitGroup() {
	for i := 1; i < 10; i++ {
		fmt.Printf("Hello %d times\n", i)
		if i == 3 {
			time.Sleep(time.Second * 2)
		}
	}
}
func waitHello() {
	fmt.Println("wait hello...")
}

/** A better helloWaitGroup (){
for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i) // Passing i as argument
	}

	time.Sleep(time.Second * 1)
}
**/
