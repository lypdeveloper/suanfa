package main

import (
	"fmt"
)

type ListNode struct{
	   Val int
	   Next *ListNode
	}
func printListFromTailToHead1( head *ListNode ) []int {
	// write code here
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	tmp := printListFromTailToHead(head.Next)
	ans = append( tmp, head.Val)
	return ans
}

func reverse(head *ListNode) *ListNode {
	var last *ListNode
	cur := head
	next := cur.Next
	for cur != nil {
		next = cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}

func printListFromTailToHead( head *ListNode ) []int {
	if head == nil {
		return []int{}
	}
	head = reverse(head)
	ret := make([]int,0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

