package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const workerCount = 3
const jobCount = 10

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d received cancellation signal\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: job channel closed\n", id)
				return
			}
			fmt.Printf("Worker %d started job %d\n", id, job)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d finished job %d\n", id, job)
		}
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// Send jobs
	go func() {
		for j := 1; j <= jobCount; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Main <--> Canceling context...")
		cancel()
	})

	wg.Wait()
	fmt.Println("Main <--> All workers done.")
}
