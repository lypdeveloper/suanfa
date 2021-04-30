package main

import (
	"fmt"
	"sync"
	"time"
)

func workingPool(jobc <-chan int, retc chan<- int, num int) {

	swg := sync.WaitGroup{}
	for i:= 0; i< num; i++ {
		swg.Add(1)
		go worker(jobc, retc, &swg)
	}
	go func() {
		swg.Wait()
		close(retc)

	}()

}

func worker(jobc <-chan int, retc chan<- int, swg *sync.WaitGroup) {
	defer swg.Done()

	for {
		t := time.NewTimer(time.Second * 3)
		select {
		case v, ok:= <-jobc:
			if !ok {
				jobc = nil
				break
			}
			retc<-v*v
		case <-t.C:
			fmt.Println("timeout")
			return
		}
	}
	//for v := range jobc {
	//	retc<-v*v
	//}
}

func gen(jobc chan<- int, numb... int) {
	go func() {
		defer close(jobc)
		for _, v := range numb {
			jobc<-v
		}
	}()

}


func main() {
	jobc := make(chan int)
	retc := make(chan int)
	gen(jobc, 1,2,3,4,5)
	workingPool(jobc, retc, 2)
	for v:= range retc {
		fmt.Println(v)
	}


}
