package main

import (
	"fmt"
	"math"
)

func main() {

	x := sqrt(1,2)
	fmt.Println(x)
}

func sqrt(i, x float64) float64 {
	res := (i+x/i)/2.0

	if math.Abs(i - res) < 0.0001 {
		return res
	} else {
		return sqrt(res, x)
	}
}