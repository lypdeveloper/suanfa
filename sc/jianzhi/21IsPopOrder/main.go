package main

import (
	"fmt"
)

func IsPopOrder1(pushV []int, popV []int) bool {
	// write code here
	var s = make([]int, 0)
	l := len(pushV)
	lPopV := len(popV)
	p := 0
	for i := 0; i < lPopV; i++ {
		if len(s) != 0 && s[len(s)-1] == popV[i] {
			s = s[:len(s)-1]
		} else {
			for p < l && pushV[p] != popV[i] {
				s = append(s, pushV[p])
				p++
			}
			p++
		}
	}
	if len(s) == 0 {
		return true
	}
	return false

}

func IsPopOrder(pushV []int, popV []int) bool {
	// write code here
	var s = make([]int, 0)
	l := len(pushV)
	lPopV := len(popV)
	p := 0
	for i := 0; i < l; i++ {
		s = append(s, pushV[i])
		for len(s) != 0 && p < lPopV && popV[p] == s[len(s)-1] {
			s = s[:len(s)-1]
			p++
		}
	}
	if len(s) == 0 {
		return true
	}
	return false

}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
