package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,10)
	for i:= 0; i<10; i++ {
		select {
		case x:=<-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("nothing")
		}
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


