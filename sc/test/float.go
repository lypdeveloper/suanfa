package main

import (
	"fmt"
	"strconv"
)

func main() {


	a := strconv.FormatFloat(0.96666343, 'G', 4, 32)
	fmt.Println(a)

	flTotalRate, _ := strconv.ParseFloat("0.96666343",64)
	fmt.Println(flTotalRate)

}
