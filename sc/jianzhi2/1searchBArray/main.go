package main

import (
	"fmt"
)

//在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//[
//[1,2,8,9],
//[2,4,9,12],
//[4,7,10,13],
//[6,8,11,15]
//]
//给定 target = 7，返回 true。
//
//给定 target = 3，返回 false。


func Find( target int ,  array [][]int ) bool {
	i, j  := 0, len(array[0])-1
	for i < len(array) && j >= 0 {
		if array[i][j] == target {
			return true
		} else if  array[i][j] < target {
			i++
		} else if array[i][j] > target {
			j--
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

