package main

import (
	"fmt"

)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var pre *ListNode
	var post *ListNode
	for {
		if pHead == nil {
			break
		}
		post = pHead.Next
		pHead.Next = pre
		pre = pHead
		pHead = post


	}
	return pre

}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
