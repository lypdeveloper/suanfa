package main

import (
	"fmt"
	"strings"
)


func FirstNotRepeatingChar( str string ) int {
	// write code here
	once := make(map[int32]int)
	for k, v := range str {

		if _, ok := once[v]; !ok && strings.LastIndex(str, string(v)) == k {
			return k
		}
		once[v] = 1
	}
	return -1

}

func main() {
	x := FirstNotRepeatingChar("google")
	fmt.Println(x)
}

