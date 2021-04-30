package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func Read() {
	lock.Lock()
	time.Sleep(time.Millisecond)
	lock.Unlock()
	wg.Done()
}


func write() {
	lock.Lock()
	x = x+1
	time.Sleep(5* time.Millisecond)
	lock.Unlock()
	wg.Done()
}
func main() {

	start := time.Now()
	for i:= 0; i< 1000; i++ {
		wg.Add(1)
		go Read()
	}


	for i:= 0; i< 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

