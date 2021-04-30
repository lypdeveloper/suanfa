package main

import (
	"fmt"
)

 type ListNode struct{
	   Val int
	   Next *ListNode
	 }
func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
	// write code here
	cur := pHead
	fast := pHead
	if pHead == nil {
		return nil
	}
	for i:= 0; i<k; i++ {
		fast = fast.Next

	}
	for fast != nil {
		fast = fast.Next
		cur = cur.Next
	}
	return cur


}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

