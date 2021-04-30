package main

import (
	"fmt"
)

func FindNumbersWithSum(array []int, sum int) []int {
	// write code here
	length := len(array)
	i, j := 0, length-1
	for i < j {
		if array[i]+array[j] == sum {
			return []int{array[i], array[j]}
		}
		if array[i]+array[j] < sum {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
