package main

import "fmt"

func partition (arr []int, start int, end int) int {
	tag := arr[end]
	i, j := start, end-1
	for i <= j {
		for i<=j && arr[i] <= tag {
			i++
		}
		for i<=j && arr[j] > tag {
			j--
		}
		if i < j {
			fmt.Println(i, j, arr)
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	fmt.Println(arr)
	fmt.Println(i, j)
	if i<=j {

	}
	arr[i], arr[end] = arr[end], arr[i]

	fmt.Println(arr)
	return i
}

func quickSort (arr []int, start int , end int) {
	if start >= end {
		return
	}
	mid := partition(arr, start, end)
	quickSort(arr, start, mid-1)
	quickSort(arr, mid+1, end)
}

func main() {
	arr := []int{1, 2, 3, 2, 9, 8, 3, 2, 0, 100, 2, 1, -1}
	//arr = []int{8, 6 }
	quickSort(arr, 0, len(arr)-1)
	//partition(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
