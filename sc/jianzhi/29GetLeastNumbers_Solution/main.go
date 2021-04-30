package main

import (
	//"container/heap"
	"fmt"
	"sort"
)
func GetLeastNumbers_Solution1( input []int ,  k int ) []int {
	// write code here
	if k > len(input) {
		return []int{}
	}
	sort.Ints(input)
	return input[0:k]
}

type Priority struct {

}
//
//func (p Priority)Len() int {
//
//}
//func (p Priority) Less(i, j int) bool {
//
//}
//
//func (p Priority) Swap(i, j int) {
//
//}


//func GetLeastNumbers_Solution( input []int ,  k int ) []int {
//	// write code here
//	heap.Init()
//}
func partition(sortArray []int, left, right int) int {
	key := sortArray[right]
	i := left - 1

	for j := left; j < right; j++ {
		if sortArray[j] <= key {
			i++
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}
	}
	sortArray[i+1], sortArray[right] = sortArray[right], sortArray[i+1]

	return i + 1
}
func GetLeastNumbers_Solution( input []int ,  k int ) []int {
	// write code here
	if k > len(input) {
		return []int{}
	}
	index := 0
	l := 0
	r := len(input) -1
	for index != k && l <r{
		index = partition(input,l, r)
		if index>k {
			r = index -1
		} else {
			l = index + 1
		}
	}
	//sort.Ints(input)
	sort.Ints(input[0:k])
	return input[0:k]
}


func main() {
	a := GetLeastNumbers_Solution([]int{1,2,3,4,5,6}, 3)
	fmt.Println(a)
}

