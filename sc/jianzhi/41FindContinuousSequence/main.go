package main

import (
	"fmt"
)

func FindContinuousSequence(sum int) [][]int {
	// write code here
	i, j := 1, 2
	curSum := i + j
	ans := [][]int{}
	for i < j && j < sum {
		if curSum == sum {
			tmp := []int{}
			for s := i; s <= j; s++ {
				tmp = append(tmp, s)
			}
			ans = append(ans, tmp)
			curSum -= i
			i++
		} else if curSum < sum {

			j++
			curSum += j
		} else if curSum > sum {
			curSum -= i
			i++

		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	FindContinuousSequence(0)

	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
}
