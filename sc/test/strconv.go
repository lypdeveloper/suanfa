package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num2 float64 =6.66
	var str string
	str = strconv.FormatFloat(num2,'f',2,32)
	fmt.Println(str)

}
