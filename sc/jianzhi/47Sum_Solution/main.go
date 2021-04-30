package main

import (
	"fmt"
)
func Sum_Solution( n int ) int {
	// write code here
	var s int
	var sum func(n int) bool
	sum = func(n int) bool {
		s += n
		return n >0 && sum(n-1)
	}
	sum(n)
	return s
}

func main() {

	fmt.Println(Sum_Solution(5))
}

