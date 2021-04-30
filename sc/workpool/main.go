package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int,100)
	results := make(chan int,100)

	for j:=0; j<2 ;j++ {
		go worker(j, jobs, results)
	}

	for i:=0; i<10; i++ {
		jobs <- i
	}
	close(jobs)

	for i:=0; i<10; i++ {
		fmt.Println(<-results)
	}


}

func worker(id int, jobs<-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("work : %d, start job : %d \n", id, job)
		results<-job * job
		time.Sleep(time.Second)
		fmt.Printf("work : %d, stop job : %d \n", id, job)
	}
}


