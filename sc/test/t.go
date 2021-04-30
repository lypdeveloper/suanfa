package main

import "fmt"

func InArray(needle interface{}, needleArr []interface{}) bool {
	for _, v := range needleArr {
		if needle == v {
			return true
		}
	}
	return false
}

func main() {
	r := InArray(12, []interface{}{1,2,3})
	fmt.Println(r)
}
