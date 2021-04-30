package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	pstack := []*TreeNode{}
	ppath := []*TreeNode{}
	qstack := []*TreeNode{}
	qpath := []*TreeNode{}
	dfs(root, p, &pstack, &ppath)
	printLine(&ppath)
	fmt.Println("**ppath", ppath)
	dfs(root, q, &qstack, &qpath)
	printLine(&qpath)
	fmt.Println("**qpath", qpath)
	i := 0
	var ans *TreeNode
	for i < len(ppath) && i < len(qpath) {

		if ppath[i].Val == qpath[i].Val {
			ans = ppath[i]
		}
		i++
	}
	return ans
}

func dfs(root, aim *TreeNode, stack *[]*TreeNode, path *[]*TreeNode) {
	printLine(path)
	if root == nil {
		return
	}
	*stack = append(*stack, root)
	if root.Val == aim.Val {
		*path = *stack
		return
	}

	dfs(root.Left, aim, stack, path)
	dfs(root.Right, aim, stack, path)
	*stack = (*stack)[:len(*stack)-1]
}

func printLine(path *[]*TreeNode) {
	for _, node := range *path {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println()
}

func DeSerialization(arr *[]int) *TreeNode {
	cur := (*arr)[0]
	*arr = (*arr)[1:]
	if cur == -99 {
		return nil
	}
	root := &TreeNode{Val: cur}
	root.Left = DeSerialization(arr)
	root.Right = DeSerialization(arr)
	return root
}

func main() {
	//t1 := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, 6, -1, -1, 7, -1, -1}
	t2 := []int{-1,0,-2,8,-99,-99,4,-99,-99,3,-99,-99}
	//n1 := TreeNode{8, nil, nil}
	//n2 := TreeNode{4, nil, nil}
	t := DeSerialization(&t2)
	fmt.Println(t)
	//pstack := []*TreeNode{}
	//ppath := []*TreeNode{}
	////dfs(t, &TreeNode{4, nil, nil}, &pstack, &ppath)
	//dfs(t, &TreeNode{5, nil, nil}, &pstack, &ppath)
	//fmt.Println(ppath)
	//ans := lowestCommonAncestor(t, &n1, &n2)
	//fmt.Println(ans)
}
