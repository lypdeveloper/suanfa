package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func find(pRootOfTree *TreeNode, lastNode *TreeNode) *TreeNode {
//
//	if pRootOfTree == nil {
//		return lastNode
//	}
//	if pRootOfTree.Left != nil {
//		lastNode = find(pRootOfTree.Left, lastNode)
//	}
//	pRootOfTree.Left = lastNode
//	lastNode.Right = pRootOfTree
//	if pRootOfTree.Right != nil {
//		lastNode = find(pRootOfTree.Ri, lastNode)
//	}
//
//
//}
//
//func Convert(pRootOfTree *TreeNode) *TreeNode {
//	// write code here
//
//}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
