package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balance(pRoot *TreeNode) (float64, bool) {
	if pRoot == nil {
		return 0, true
	}
	ld, lbool := balance(pRoot.Left)
	rd, rbool := balance(pRoot.Right)
	abs := math.Abs(ld-rd)
	if lbool && rbool &&  abs<= 1 {
		return math.Max(ld, rd) + 1, true
	} else {
		return 0, false
	}
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	_, bo := balance(pRoot)

	return bo
}


func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
