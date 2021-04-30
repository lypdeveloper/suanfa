package main

import (


)


func canJump(nums []int) bool {
	jumps := []int{}
	for i := 0; i < len(nums); i++ {
		cur := i + nums[i]
		jumps = append(jumps, cur)
	}
	maxjump := 0
	for i:= 0; i <len(nums) ; i++ {
		if i > maxjump {
			//break
			return false
		}
		if maxjump<jumps[i] {
			maxjump = jumps[i]
		}

	}
	return true
}
func main() {


}
