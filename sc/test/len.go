package main

import "fmt"

func main() {

	c := make(chan int, 10)
	c<-1
	c<-1
	c<-1
	c<-1
	c<-1
	c<-1
	<-c
	fmt.Println(len(c))
}

