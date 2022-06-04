package main

import (
	"fmt"
	"math/rand"
	"time"
)

func workerpoolstest() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workerJob(w, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for ja := 0; ja < numJobs; ja++ {
		<-results
	}
}

func workerJob(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker id", j, "started job", j)
		ll := time.Duration(rand.Intn(10))
		time.Sleep(time.Second * ll)
		fmt.Println("worker id", j, "finished job", j, "in", ll, "seconds")
		results <- j * 2
	}
}
