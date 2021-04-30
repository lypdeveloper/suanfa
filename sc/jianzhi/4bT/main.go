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
func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
	// write code here
	if len(pre) <= 0 || len(vin) <= 0 {
		return nil
	}
	kroot := 0
	root := pre[0]
	for k, v := range vin {
		if v == root {
			kroot = k
		}
	}
	rootnode := &TreeNode{root, nil, nil}
	rootnode.Left = reConstructBinaryTree(pre[1:kroot+1], vin[:kroot])
	rootnode.Right = reConstructBinaryTree(pre[kroot+1:], vin[kroot+1:])
	return rootnode
}

func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

