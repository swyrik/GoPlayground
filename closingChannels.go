package main

import "fmt"

func closingChannelsTest() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recevied job", j)

			} else {
				fmt.Println("jobs receving completed")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}

	close(jobs)
	fmt.Println("all jobs are sent")
	<-done

}
