package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func symme(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1== nil && pRoot2 == nil {
		return true
	}
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if pRoot1.Val == pRoot2.Val && symme(pRoot1.Left, pRoot2.Right) && symme(pRoot1.Right, pRoot2.Left) {
		return true
	} else {
		return false
	}
}

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	return symme(pRoot, pRoot)
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
