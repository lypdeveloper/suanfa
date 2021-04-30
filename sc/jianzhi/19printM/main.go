package main

import "fmt"

func printMatrix1( matrix [][]int , li, ri, lj,rj int) []int {
	// write code here
	ans := make([]int, 0)
	if li == ri {
		ans = append(ans, matrix[lj][li])
		return ans
	}


	i := li
	j := lj
	for ; i<ri; i++{
		ans = append(ans, matrix[j][i])
	}
	i = ri
	j = lj
	for ; j<rj; j++{
		ans = append(ans, matrix[j][i])
	}
	i = ri
	j=rj
	for ; i>li; i-- {
		ans = append(ans, matrix[j][i])
	}
	j = rj
	i = li
	for ; j>lj; j-- {
		ans = append(ans, matrix[j][i])
	}
	return ans
}
func printMatrix( matrix [][]int ) []int {

	length := len(matrix)
	ans := make([]int,0)
	li := 0
	ri := length - 1
	lj := 0
	rj := length -1
	for length>0 {

		ret := printMatrix1(matrix, li,ri,lj, rj)
		ans = append(ans, ret ...)
		li++
		ri--
		lj++
		rj--
		length -= 2
	}
	return ans
}

func main() {
	//a := [][]int{[]int{1,2,3,0},[]int{4,5,6,0},[]int{7,8,9,0},{10,11,12,0}}
	a := [][]int{{4}}
	fmt.Println(printMatrix(a))
}

