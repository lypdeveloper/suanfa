package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, expectNumber int, curPath []int, ans *[][]int) {
	if root == nil {
		return
	}
	curPath = append(curPath, root.Val)
	if root.Val == expectNumber && root.Left == nil && root.Right == nil{
		*ans = append(*ans, curPath)
	}


	if root.Left != nil {
		dfs(root.Left, expectNumber-root.Val, curPath, ans)
	}
	if root.Right != nil {
		dfs(root.Right, expectNumber-root.Val, curPath, ans)
	}
	curPath = (curPath)[:len(curPath)-1]
}

func FindPath(root *TreeNode, expectNumber int) [][]int {
	// write code here
	curPath := make([]int, 0)
	ans := make([][]int, 0)
	dfs(root, expectNumber, curPath, &ans)
	return ans

}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
