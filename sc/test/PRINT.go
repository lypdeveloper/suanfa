package main

import (
	"fmt"
)
func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)

	a := [][]int{
		[]int{1,2,3,4,5},
		[]int{6,7,8,9,10},
		[]int{11,12,13,14,15},
		[]int{16,17,18,19,20},
		[]int{21,22,23,24,25},

	}
	loopPrint(a)
}

func loopPrint(a [][]int) {
	lenOne := len(a)
	lenTwo := len(a[0])
	printArr := make([]int, 0)
	startOne := 0
	startTwo := 0
	j := 0
	i := 0
	for startOne < lenOne && startTwo < lenTwo {
		j = startTwo
		for i = startOne; i<lenTwo-1; i++ {
			fmt.Println(1, a[j][i])
			printArr = append(printArr, a[j][i])
		}

		for ; j<lenOne -1; j++ {
			fmt.Println(2, a[j][i])
			printArr = append(printArr, a[j][i])
		}
		for ; i>startOne; i-- {
			fmt.Println(3, a[j][i])
			printArr = append(printArr, a[j][i])
		}
		for ; j>startTwo; j-- {
			fmt.Println(4, a[j][i])
			printArr = append(printArr, a[j][i])
		}
		startOne++
		startTwo++
		lenOne--
		lenTwo--
		fmt.Printf("tage %d,%d,%d,%d\n", startOne, startTwo, lenOne, lenTwo)

	}
	for _, v := range printArr {
		fmt.Printf("%d ", v)
	}


	return

}