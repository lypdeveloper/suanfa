package main

import (
	"fmt"
)

 type ListNode struct{
	   Val int
	   Next *ListNode
	 }
func FindKthToTail1( pListHead *ListNode ,  k int ) *ListNode {
	// write code here
	f,s := pListHead,pListHead
	for i:=0; i< k ;i++ {
		if f == nil {
			return nil
		}
		f = f.Next
	}
	for f!=nil {
		f = f.Next
		s = s.Next
	}
	return s


}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

