package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	rwMux := new(sync.RWMutex)
	names := make(map[int]string)
	for i := 0; i < 10; i++ {
		names[i] = fmt.Sprintf("name %d",i)
	}

	go func() {
		for i := 0; i < 1000; i++ {
			go func(i int) {
				//rwMux.RLock()
				//time.Sleep(time.Second)
				fmt.Println(names[1],i)
				//rwMux.RUnlock()
			}(i)
		}
	}()

	time.Sleep(500)
	go func() {
		for i := 0; i < 10; i++ {
			go func(i int) {
				rwMux.Lock()
				names[i] = fmt.Sprintf("hello %d",i)
				fmt.Println("writer....")
				rwMux.Unlock()
			}(i)
		}
	}()
	time.Sleep(time.Second * 10)
}
