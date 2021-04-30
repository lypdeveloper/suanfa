package main

import "fmt"

func jumpFloor( number int ) int {
	// write code here
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	return jumpFloor(number-1) + jumpFloor(number-2)
}

func jumpFloor1( number int ) int {
	// write code here

	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	p1 := 1
	p2 := 2
	p3 := 0
	for i :=3;i<=number;i++ {
		p3 = p1+p2
		p1 = p2
		p2 = p3

	}
	return p3
}


func main() {
	fmt.Println(jumpFloor1(4))
}

