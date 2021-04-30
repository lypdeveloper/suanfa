package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge1(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	head := &ListNode{0, nil}
	ans := head
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			head.Next = &ListNode{pHead1.Val, nil}
			head = head.Next
			pHead1 = pHead1.Next
		} else {
			head.Next = &ListNode{pHead2.Val, nil}
			head = head.Next
			pHead2 = pHead2.Next
		}

	}
	if pHead1 != nil {
		for pHead1 != nil {
			head.Next = &ListNode{pHead1.Val, nil}
			head = head.Next
			pHead1 = pHead1.Next
		}
	}
	if pHead2 != nil {
		for pHead2 != nil {
			head.Next = &ListNode{pHead2.Val, nil}
			head = head.Next
			pHead2 = pHead2.Next
		}
	}
	return ans.Next
}


func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	ans := &ListNode{}
	cur := ans
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val > pHead2.Val {
			cur.Next = pHead2
			pHead2 = pHead2.Next
			cur = cur.Next
			continue
		}
		if pHead1.Val <= pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
			cur = cur.Next
			continue
		}
	}

	if pHead1 != nil {
		cur.Next = pHead1
	}
	if pHead2 != nil {
		cur.Next = pHead2
	}
	return ans.Next

}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
