package main

import (
	"fmt"
)


 type TreeNode struct {
	   Val int
	   Left *TreeNode
	   Right *TreeNode
	 }
func Mirror1( pRoot *TreeNode ) *TreeNode {
	// write code here
	//后序遍历

	if pRoot == nil {
		return nil
	}
	l := Mirror(pRoot.Left)
	r := Mirror(pRoot.Right)
	pRoot.Left = r
	pRoot.Right = l
	return pRoot
}

func Mirror( pRoot *TreeNode ) *TreeNode {
	// write code here
	//后序遍历

	if pRoot == nil {
		return nil
	}
	q := make([]*TreeNode, 0)
	q = append(q, pRoot)
	for len(q)>0 {
		len := len(q)
		for len >0 {
			len--
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q,node.Left)
			}
			if node.Right != nil {
				q = append(q,node.Right)
			}
			tmp := node.Left
			node.Left = node.Right
			node.Right = tmp

		}
	}

	return pRoot
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
