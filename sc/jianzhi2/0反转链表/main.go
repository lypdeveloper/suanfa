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
	cur := pHead
	var head *ListNode
	last := head
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}

	return last

}

func printLink(pHead *ListNode) {
	for pHead != nil {
		fmt.Printf("% p  %d  %p\n", pHead, pHead.Val, pHead.Next)
		pHead = pHead.Next
	}
}


func main() {
	n0 := &ListNode{0,nil}
	n00 := &ListNode{0,nil}
	n000 := &ListNode{0,nil}
	n1 := &ListNode{1,nil}
	n11 := &ListNode{1,nil}
	n2 := &ListNode{2,nil}
	n0.Next=n00
	n00.Next=n000
	n000.Next = n1
	n1.Next=n11
	n11.Next=n2
	n2.Next=nil
	//printLink(n0)
	p := ReverseList(n0)
	printLink(p)



}
