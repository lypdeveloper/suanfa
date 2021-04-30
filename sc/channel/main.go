package main

import (
	"fmt"
)



func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

