package main

import (
	"fmt"
	"time"
)

func closeCh(ch chan int) {
	close(ch)
}

func main() {
	t := time.NewTimer(time.Second)
	ch := make(chan int, 10)
	go func() {
		for {
			select {
			case <-t.C:
				close(ch)
			case <-ch:
				fmt.Println("close")
				return
			}
		}
	}()
	time.Sleep(2*time.Second)
}
