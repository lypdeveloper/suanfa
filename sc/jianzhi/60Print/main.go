package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Print(pRoot *TreeNode) [][]int {
	// write code here
	if pRoot == nil {
		return [][]int{}
	}
	q := make([]*TreeNode, 0)
	tmp := make([]int, 0)
	ans := make([][]int, 0)
	q = append(q, pRoot)
	for len(q) > 0 {
		ql := len(q)
		for i := 0; i< ql; i++ {
			pop := q[0]
			q = q[1:]
			tmp = append(tmp, pop.Val)
			if pop.Left != nil {
				q = append(q, pop.Left)
			}
			if pop.Right != nil {
				q = append(q, pop.Right)
			}
		}
		ans = append(ans, tmp)
		tmp = []int{}
	}
	return ans
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
