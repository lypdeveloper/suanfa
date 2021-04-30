package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	//生产者
	go func() {
		defer func() {
			close(c)
			wg.Done()
		}()
		for i := 0; i < 10; i++ {
			c <- rand.Int()
		}

	}()
	time.Sleep(1 * time.Second)
	//消费者
	ti := time.After(1*time.Second)
	go func() {
		defer wg.Done()
		//for i := 0; i < 10; i++ {
		//	rev := <-c
		//	fmt.Println(rev)
		//}
		//for val := range c {
		//	fmt.Println(val)
		//}
		for  {
			select {
			case val, ok:= <-c:
				if ok {
					fmt.Println(val)
				}
			case <-ti:
				fmt.Println("timeout")
				return

			}


		}

	}()

	wg.Wait()

}
