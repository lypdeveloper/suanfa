package main

import (
	"fmt"
	"sync"
)

func Produce(num ...int) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for v := range num {
			c <- v
		}
	}()
	return c
}

func Square(in <-chan int) <-chan int {
	o := make(chan int)
	go func() {
		defer close(o)
		for v := range in {
			o <- v * v
		}
	}()
	return o
}

var wg1 = sync.WaitGroup{}
func Merge(ins ...<-chan int) <-chan int {

	o := make(chan int)

	coll := func(in <-chan int) {
		defer wg1.Done()
		for v := range in {
			o <- v
		}
	}

	for _, v := range ins {
		wg1.Add(1)
		go coll(v)
	}


	go func() {
		wg1.Wait()
		close(o)
	}()

	return o
}

func main() {
	in := Produce(1, 2, 3, 4, 5)
	o1 := Square(in)
	o2 := Square(in)
	o3 := Square(in)
	o := Merge(o1, o2, o3)

	//for i:=0;i<3;i++ {
	//	fmt.Println(<-o)
	//}
	for v := range o {
		fmt.Println(v)
	}
}
