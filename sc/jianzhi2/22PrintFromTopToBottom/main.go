package main

import (
	"fmt"
)


  type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }

func PrintFromTopToBottom( root *TreeNode ) []int {
	// write code here
	if root == nil {
		return []int{}
	}
	q := make([]*TreeNode, 0)
	ans := make([]int, 0)
	q = append(q, root)
	for len(q) > 0 {
		cur := q[0]
		ans = append(ans, cur.Val)
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
		q = q[1:]

	}
	return ans

}

func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

