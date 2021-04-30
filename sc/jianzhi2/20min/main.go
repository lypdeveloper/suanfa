package main

import "fmt"

var stack = make([]int,0)
var minStack = make([]int,0)

func Push(node int) {
	// write code here
	stack = append(stack, node)
	lenMin := len(minStack)
	if len(minStack) != 0 {
		if minStack[lenMin -1] > node {
			minStack = append(minStack, node)
		} else {
			minStack = append(minStack, minStack[lenMin -1])
		}
	} else {
		minStack = append(minStack, node)
	}
}
func Pop() {
	// write code here
	l := len(stack)
	stack = stack[:l-1]
	minStack = minStack[:l-1]
}
func Top() int {
	// write code here
	l := len(stack)
	return stack[l-1]
}
func Min() int {
	// write code here
	l := len(minStack)
	return minStack[l-1]
}

func main() {
	Push(2)
	Push(3)
Push(1)

Push(1)

fmt.Println(stack)
fmt.Println(minStack)

}

