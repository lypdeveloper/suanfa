package main

import (
	"fmt"
)
func NumberOf1Between1AndN_Solution( n int ) int {
	// write code here
	return 0
}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

