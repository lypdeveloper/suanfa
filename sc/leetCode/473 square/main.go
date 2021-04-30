package main

import (
	"fmt"
	"sort"
)


func makesquare(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] <= nums[j] {
			return false
		}
		return true
	})
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	blen := sum / 4
	if sum%4 != 0 {
		return false
	}

	buket := make([]int, 4)
	return put(0, nums, buket, blen)


}

func put(i int, nums []int, buket []int, blen int, ) bool{
 	if i > len(nums) {
 		return true
	}
	for j := 0; j < len(buket); j++ {
		if buket[j] + nums[i] > blen {
			continue
		}
		buket[j] += nums[i]
		if put(i+1, nums, buket, blen) {
			return true
		}
		buket[j] -= nums[i]
	}
	return false
}

func main() {
	nums := []int{3,3,5,73,1,2}

	sort.Slice(nums, func(i, j int) bool {
		if nums[i] <= nums[j] {
			return false
		}
		return true
	})
	fmt.Println(nums)

}
