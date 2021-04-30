package main

import "fmt"

func Power( base float64 ,  exponent int ) float64 {
	// write code here
	if exponent < 0 {
		base = 1/base
		exponent = -exponent
	}
	return qPower(base, exponent)


}

func qPower( base float64 ,  exponent int ) float64 {
	// write code here
	if exponent == 0 {
		return 1.0
	}
	if exponent == 1 {
		return base
	}
	if exponent/2 == 0 {
		return qPower(base, exponent/2) * qPower(base, exponent/2)
	} else {
		return qPower(base, exponent/2) * qPower(base, exponent/2) * base
	}
}


func main() {
	q := Power(2,3)
	fmt.Println(q)
}

