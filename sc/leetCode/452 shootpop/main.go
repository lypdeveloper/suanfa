package main

import (
	"fmt"
	"sort"
)


func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	//排序
	sort.SliceStable(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		}
		return false
	})
	tmpRange := []int{points[0][0], points[0][1]}
	ans := 1

	for i := 1; i < len(points); i++ {
		if points[i][0] > tmpRange[1] {
			ans++
			tmpRange[0] = points[i][0]
			tmpRange[1] = points[i][1]
		} else {
			if tmpRange[1] > points[i][1] {
				tmpRange[1] = points[i][1]
			}
		}
	}

	return ans
}

func main() {

	points := [][]int{
		{1,5},
		{8,9},
		{4,6},

	}
	shots := findMinArrowShots(points)
	fmt.Println(shots)
}
