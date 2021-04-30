package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeDepth(pRoot *TreeNode) int {
	// write code here

	if pRoot == nil {
		return 0
	}
	if pRoot.Left == nil && pRoot.Right == nil {
		return 1
	}
	leftDepth := TreeDepth(pRoot.Left)
	rightDepth := TreeDepth(pRoot.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
