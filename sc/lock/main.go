package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)
func main() {

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Print(x)

}

func add() {
	lock.Lock()
	for i:=0; i<5000; i++ {
		lock.Lock()
		x = x+1
		lock.Unlock()
	}
	lock.Unlock()
	wg.Done()
}


