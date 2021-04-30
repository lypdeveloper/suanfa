package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialization(root *TreeNode) []int {

	ans := make([]int, 0)
	ans = append(ans, root.Val)
	arrl := []int{}
	arrr := []int{}
	if root.Left != nil {
		arrl = Serialization(root.Left)
	} else {
		arrl = append(arrl, -1)
	}
	if root.Right != nil {
		arrr = Serialization(root.Right)
	} else {
		arrr = append(arrr, -1)
	}

	ans = append(ans, arrl...)
	ans = append(ans, arrr...)
	return ans
}

func Serialization1(root *TreeNode, arr *[]int) []int {
	if root == nil {
		*arr = append(*arr, -1)
		return *arr
	}
	*arr = append(*arr, root.Val)
	Serialization1(root.Left, arr)
	Serialization1(root.Right, arr)
	return *arr

}

func DeSerialization(arr *[]int) *TreeNode {
	cur := (*arr)[0]
	*arr = (*arr)[1:]
	if cur == -1 {
		return nil
	}
	root := &TreeNode{Val: cur}
	root.Left = DeSerialization(arr)
	root.Right = DeSerialization(arr)
	return root
}

func Serialize(root *TreeNode) *TreeNode {
	// write code here
	arr := make([]int, 0)
	a := Serialization1(root, &arr)
	fmt.Println(a)
	root = DeSerialization(&a)
	return root
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
