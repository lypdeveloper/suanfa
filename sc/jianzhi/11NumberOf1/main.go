package main

import (
	"fmt"
	"math"
)

func NumberOf1(n int) int {
	// write code here



	cnt := 0
	tmp := 1
	for i:=0; i<=31; i++{
		if tmp&n != 0 {
			cnt++
		}

		tmp = tmp << 1
	}


	return cnt
}


func main() {
	fmt.Printf("%-b \n", -1)
	var a int32
	//a = -2147483647
	fmt.Printf("%b \n", math.MaxInt8)
	fmt.Printf("%b\n", a)
	c := NumberOf1(-1)
	fmt.Println(c)
}
