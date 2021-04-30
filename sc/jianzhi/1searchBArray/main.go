package main

import (
	"fmt"
)
func Find( target int ,  array [][]int ) bool {
	// write code here
	n := len(array)
	if n == 0 {
		return false
	}
	m:=len(array[0])
	if m == 0 {
		return false
	}

	for i, j := 0, m-1 ; i < n -1 && j >0; {
		if target == array[i][j] {
			return true
		}
		if target > array[i][j] {
			i ++
			continue
		}
		if array[i][j] > target {
			j++
			continue
		}

	}
	return false

}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

