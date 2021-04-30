package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	head := &ListNode{-1, nil}
	head.Next = pHead

	last := head
	cur := pHead
	for cur != nil &&cur.Next != nil {
		if cur.Val == cur.Next.Val {
			needle := cur.Val
			for cur !=nil && cur.Val == needle {
				//cur.Next = cur.Next.Next
				//last = cur
				cur = cur.Next
			}
			last.Next = cur
		} else {
			last = cur
			cur = cur.Next
		}
	}
	return head.Next
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

	p := deleteDuplication(n0)
	printLink(p)


}
