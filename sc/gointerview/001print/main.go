package main

import (
	"fmt"
)

//var num, letter = make(chan int),  make(chan int)
//var str = "abcdefghigklmno"
//func printNum(wg *sync.WaitGroup) {
//
//	for i := 1; i<=14; i+=2 {
//		<-num
//		fmt.Println(i)
//		fmt.Println(i+1)
//		letter<-1
//	}
//
//}
//func printLetter(wg *sync.WaitGroup) {
//
//	for i := 1; i<=14; i+=2 {
//		<-letter
//		fmt.Println(string(str[i]))
//		fmt.Println(string(str[i+1]))
//		num<-1
//	}
//
//
//
//}
//
//
//func main() {
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//
//	go printNum(&wg)
//	go printLetter(&wg)
//	num<-1
//	wg.Wait()
//
//
//}

var chan1 = make(chan bool )
var chan2 = make(chan bool)
var index = make(chan bool)
var str = "ABCDEFGHIJ"

func func1() {
	//for i := 1; i < 11;  i += 2 {
	//	<-chan1
	//	fmt.Println(i)
	//	fmt.Println(i + 1)
	//	chan2 <- true
	//}
	i:=1
	for {
		select {
		case <-chan1:
			fmt.Println(i)
			fmt.Println(i + 1)
			i += 2
			chan2<-true
		default:
			continue
		}

	}
}
func func2() {
	size := len(str)
	for i := 0; i < size; i += 2 {
		<-chan2
		fmt.Println(str[i])
		fmt.Println(str[i+1])
		chan1 <- true
	}
	index <- true
}
func main() {
	//time.Sleep(10*time.Second)
	go func1()
	go func2()
	chan1 <- true


	<-index
}