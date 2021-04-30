package main

import (
	"fmt"
)

 type TreeNode struct {
	   Val int
	   Left *TreeNode
	   Right *TreeNode
	 }

func reConstructBinaryTree1( pre []int ,  vin []int ) *TreeNode {
	// write code here
	if len(pre) == 0 || len(vin) == 0 {
		return nil
	}
	tnode := pre[0]
	tree := &TreeNode{Val: tnode}
	for k, v := range vin {
		if v == tnode {
			tree.Left = reConstructBinaryTree(pre[1:k+1], vin[:k])
			tree.Right = reConstructBinaryTree(pre[k+1:], vin[k+1:])
			break
		}
	}
	return tree
}

//输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
	// write code here

}

func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

