package main

import (
	"fmt"
)
func reOrderArray( array []int ) []int {
	// write code here
	ans := make([]int,0)
	ji := make([]int,0)
	for _, v := range array {
		if v%2==1 {
			ji = append(ji, v)
		} else {
			ans = append(ans, v)
		}
	}
	return append(ji, ans...)
}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

