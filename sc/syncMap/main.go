package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup
var m2 sync.Map

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i , i+100)
			ret , _:= m2.Load(i)
			fmt.Println(i, " ",ret)
			wg.Done()
		}(i)
	}

	fmt.Println("ee")
	wg.Wait()
}
