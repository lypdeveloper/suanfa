package main

import "fmt"


func jumpFloorII( number int ) int {
	// write code here

	if number == 1 {
		return 1
	}

	return 2 * jumpFloorII(number-1)

}



func main() {
	fmt.Println((4))
}

