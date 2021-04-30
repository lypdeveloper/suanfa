package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go Controler(&wg, 10)
	wg.Wait()
}

func Controler(wg *sync.WaitGroup, num int) {

	ctx, cancel := context.WithCancel(context.Background())
	catch := make(chan struct{}, 0)
	dogch := make(chan struct{}, 0)
	exit := make(chan struct{}, 0)
	finish := make(chan struct{}, 0)
	go Cat(wg, ctx, catch, dogch, finish)
	go Dog(wg, ctx, dogch, exit, finish)
	sum := 0
	go func() {
		defer wg.Done()
		for {
			select {
			case <-exit:
				sum++
				if sum == num {
					close(finish)
					cancel()
					return
				}
				catch <- struct{}{}
			}
		}
	}()
	catch <- struct{}{}
}

func Cat(wg *sync.WaitGroup, ctx context.Context, catch, dogch, finish chan struct{}) {
	defer wg.Done()
	for {
		select {
		case x := <-finish:
			fmt.Println(x)
			fmt.Println("cat finish")
			//return
		case <-ctx.Done():
			fmt.Println("cat done")
			return
		case <-catch:
			fmt.Println("cat")
			dogch <- struct{}{}
		}
	}
}

func Dog(wg *sync.WaitGroup, ctx context.Context, dogch, exit, finish chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-finish:
			fmt.Println("dog finish")
			//return
		case <-ctx.Done():
			fmt.Println("dog done")
			return
		case <-dogch:
			fmt.Println("dog")
			exit <- struct{}{}
		}
	}
}
