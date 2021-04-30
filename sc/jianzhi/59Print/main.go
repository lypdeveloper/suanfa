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
type node struct {
	Val *TreeNode
	Order int
	Depth int
}

func Print_wrong(pRoot *TreeNode) [][]int {
	// write code here
	q := make([]node, 0)
	ans := make([][]int, 0)
	q = append(q, node{pRoot, 1, 1})
	line := []int{}
	curDep := 1
	for len(q) > 0 {

		pop := q[0]
		q = q[1:]

		if pop.Depth != curDep {
			curDep = pop.Depth
			ans = append(ans,line)
			line = line[:0]
		}

		line = append(line, pop.Val.Val)

		if pop.Order == 1 {
			if  pop.Val.Right != nil {
				q = append(q, node{pop.Val.Right, 0, pop.Depth+1})
			}
			if  pop.Val.Left != nil {
				q = append(q, node{pop.Val.Left, 0, pop.Depth+1})
			}
		}
		if pop.Order == 0 {
			if  pop.Val.Left != nil {
				q = append(q, node{pop.Val.Left, 1, pop.Depth+1})
			}
			if  pop.Val.Right != nil {
				q = append(q, node{pop.Val.Right, 1, pop.Depth+1})
			}
		}


	}
	return ans
}

func reverse(line []int, i, j int) {
	for i < j {
		line[i], line[j] = line[j], line[i]
		i++
		j--
	}
}

func Print(pRoot *TreeNode) [][]int {
	ans  := make([][]int, 0)
	if pRoot == nil {
		return ans
	}
	stack := make([][]*TreeNode, 2)
	cur := 1
	stack[1] = append(stack[1], pRoot)

	line := make([]int, 0)
	for len(stack[0]) != 0 || len(stack[1]) != 0 {
		pop := stack[cur][0]
		line = append(line, pop.Val)
		stack[cur] = stack[cur][1:]
		if len(stack[cur]) == 0 {
			//if cur == 0 {
			//	reverse(line, 0, len(line)-1)
			//}
			ans = append(ans, line)
			line = []int{}

		}



		if cur == 1 {
			if pop.Right != nil {
				stack[0] = append(stack[0], pop.Right)
			}
			if pop.Left != nil {
				stack[0] = append(stack[0], pop.Left)
			}
		}
		if cur == 0 {
			if pop.Left != nil {
				stack[1] = append(stack[1], pop.Left)
			}
			if pop.Right != nil {
				stack[1] = append(stack[1], pop.Right)
			}



		}
		if len(stack[cur]) == 0 {

			cur = int(math.Abs(float64(1-cur)))
		}



	}
	return ans
}

func Print2(pRoot *TreeNode) [][]int {
	ans  := make([][]int, 0)
	if pRoot == nil {
		return ans
	}
	stack := make([][]*TreeNode, 2)
	cur := 1
	stack[1] = append(stack[1], pRoot)

	line := make([]int, 0)
	for len(stack[0]) != 0 || len(stack[1]) != 0 {
		sl := len(stack[cur])
		pop := stack[cur][sl-1]
		line = append(line, pop.Val)
		stack[cur] = stack[cur][:sl-1]
		if len(stack[cur]) == 0 {
			//if cur == 0 {
			//	reverse(line, 0, len(line)-1)
			//}
			ans = append(ans, line)
			line = []int{}

		}



		if cur == 1 {
			if pop.Right != nil {
				stack[0] = append(stack[0], pop.Right)
			}
			if pop.Left != nil {
				stack[0] = append(stack[0], pop.Left)
			}
		}
		if cur == 0 {
			if pop.Left != nil {
				stack[1] = append(stack[1], pop.Left)
			}
			if pop.Right != nil {
				stack[1] = append(stack[1], pop.Right)
			}



		}
		if len(stack[cur]) == 0 {

			cur = int(math.Abs(float64(1-cur)))
		}



	}
	return ans
}


func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
