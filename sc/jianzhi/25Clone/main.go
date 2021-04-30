package main

import (
	"fmt"
)

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

func Clone(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	//write your code here
	headCopy := head
	for headCopy != nil {
		copyMy := RandomListNode{Label: headCopy.Label}
		copyMy.Next = headCopy.Next
		headCopy.Next = &copyMy
		headCopy = copyMy.Next
	}
	headCopy2 := head

	for  headCopy2 != nil{
		if headCopy2.Random != nil {
			headCopy2.Next.Random = headCopy2.Random.Next
		}

		headCopy2 = headCopy2.Next.Next


	}
	p := head
	newHead := p.Next
	np := newHead
	for np.Next != nil {
		p.Next = p.Next.Next
		p = p.Next
		np.Next = np.Next.Next
		np = np.Next
	}
	p.Next = nil

	return newHead
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
