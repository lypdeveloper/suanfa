package main

import (
	"fmt"
	"math"
)

func main() {
	ch1 := make(chan int,100)
	ch2 := make(chan int,100)

	go f1(ch1)
	go f2(ch1, ch2)
	//for ret := range ch2 {
	//	fmt.Println(ret)
	//}

	for i:=0; i<100; i++ {
		q := <-ch2
		fmt.Println(q)
	}

}

func f1(ch  chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- int(math.Pow(float64(tmp), 2))
	}
	close(ch2)
}

