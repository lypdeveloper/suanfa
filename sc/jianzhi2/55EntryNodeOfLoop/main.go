package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	if pHead == nil {
		return nil
	}
	fast := pHead
	slow := pHead
	for fast != nil && fast.Next != nil{
		if fast == slow {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	bang := fast
	cnt := 0
	for {
		cnt++
		fast = fast.Next
		if fast == bang {
			break
		}
	}
	fast = pHead
	slow = pHead
	for i:=0; i< cnt; i++ {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
