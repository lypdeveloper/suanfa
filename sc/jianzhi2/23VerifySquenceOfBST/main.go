package main

import (
	"fmt"

)
func VerifySquenceOfBST( sequence []int ) bool {
	// write code here
	//终止条件
	if len(sequence) == 0 {
		return false
	}
	if len(sequence) == 2 {
		return true
	}

	ha := len(sequence)-1
	node := sequence[len(sequence)-1]
	for k, v := range sequence {
		if v > node {
			ha = k
			break
		}
	}
	if ha != -1 {
		for i := ha; i<len(sequence)-1 ; i++ {
			if sequence[i] < node {
				return false
			}
		}


	}

	left := true
	right := true
	if len(sequence[:ha]) > 0 {
		left = VerifySquenceOfBST(sequence[:ha])
	}
	if len(sequence[ha:len(sequence)-1]) > 0  {
		right = VerifySquenceOfBST(sequence[ha:len(sequence)-1])
	}


	return left && right

}


func main() {
	data:= []int{4,8,6,12,16,14,10}
	data= []int{7,4,6,5}
	data= []int{1,2,3,4,5}
	x := VerifySquenceOfBST(data)
	fmt.Println(x)
}

