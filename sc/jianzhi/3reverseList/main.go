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

func printListFromTailToHead( head *ListNode ) []int {
	// write code here
	if head == nil {
		return []int{}
	}
	printed := printListFromTailToHead(head.Next)
	printed = append(printed, head.Val)
	return printed
}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

