package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrder(pRoot *TreeNode, arr*[]*TreeNode)  {
	if pRoot.Left != nil {
		InOrder(pRoot.Left, arr)
	}
	*arr = append(*arr, pRoot)
	if pRoot.Right != nil {
		InOrder(pRoot.Right, arr)
	}
}
func KthNode( pRoot *TreeNode ,  k int ) *TreeNode {
	// write code here
	arr := make([]*TreeNode, 0)
	InOrder(pRoot, &arr)
	return arr[k-1]


}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
