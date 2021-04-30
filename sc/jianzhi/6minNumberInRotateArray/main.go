package main

import (
	"fmt"
	//"sort"
)

func minNumberInRotateArray1(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 0 {
		return 0
	}
	l, r := 0, len(rotateArray)-1
	for l <= r {
		mid := (l + r) / 2
		if rotateArray[mid-1] > rotateArray[mid] {
			return rotateArray[mid]
		}
		if rotateArray[mid] >= rotateArray[0] {
			l = mid + 1
		} else {
			r = mid
		}
		if l == r {
			if rotateArray[l-1] > rotateArray[l] {
				return rotateArray[l]
			}
		}
	}
	return 0
}

func minNumberInRotateArray(rotateArray []int) int {
	if len(rotateArray) == 0 {
		return 0
	}
	i := 0
	j := len(rotateArray) - 1
	for i < j {

		mid := i + (j-i)/2
		if j-i == 1 && rotateArray[i] > rotateArray[j] {
			return rotateArray[j]
		}
		if rotateArray[i] <= rotateArray[mid] {
			i = mid
		} else {
			j = mid
		}
	}
	return 0
}

func findNumInRotateArray(rotateArray []int, aim int) int {
	if len(rotateArray) == 0 {
		return 0
	}
	i := 0
	j := len(rotateArray) - 1
	for i <= j {
		mid := i + (j-i)/2
		if rotateArray[mid] == aim {
			return mid
		}
		if rotateArray[i] < rotateArray[mid] {
			if aim >= rotateArray[i] && aim <= rotateArray[mid] {
				j = mid - 1
				continue
			} else {
				i = mid + 1
				continue
			}
		} else {
			if aim >= rotateArray[mid] && aim <= rotateArray[j] {
				i = mid + 1
				continue
			} else {
				j = mid - 1
				continue
			}
		}
	}
	return -1
}

func main() {
	arr := []int{3, 4, 5, 1, 2}
	r := minNumberInRotateArray(arr)
	fmt.Println(r)

	ans := findNumInRotateArray(arr, 1)
	fmt.Println(ans)
}
