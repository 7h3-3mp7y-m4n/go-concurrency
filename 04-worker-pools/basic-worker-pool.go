package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerCount = 3
	jobCount    = 5
)

func main() {
	test()
	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= workerCount; i++ {
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for j := 0; j < jobCount; j++ {
		wg.Add(1) // One wg.Add per job
		jobs <- j
	}

	close(jobs) // Tell workers there are no more jobs
	wg.Wait()   // Wait for all jobs to finish
	fmt.Println("All workers done!")
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		wg.Done() // Signal this job is done
	}
}

func test() {
	newwork := 10
	newJobs := 20

	jobies := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= newwork; i++ {
		go worker(i, jobies, &wg)
	}
	for j := 1; j <= newJobs; j++ {
		wg.Add(1)
		jobies <- j
	}
	close(jobies)
	wg.Wait()
	fmt.Println("the second version is done!")
}
