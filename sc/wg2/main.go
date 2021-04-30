package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("hhh", i)
			wg.Done()
		}(i)
	}

	fmt.Println("ee")
	wg.Wait()
}

func hello(i int) {
	fmt.Println("aaaa", i)
	wg.Done()

}