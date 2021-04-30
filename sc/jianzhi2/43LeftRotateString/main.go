package main

import (
	"fmt"
)

func reStr(str []byte, i, j int) {
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}
func LeftRotateString(str string, n int) string {
	// write code here
	if len(str) == 0 {
		return ""
	}
	n = n%9
	arrStr := []byte(str)
	reStr(arrStr, 0, n-1)
	reStr(arrStr, n, len(arrStr) -1)
	reStr(arrStr, 0, len(arrStr) -1)

	return string(arrStr)
}

func main() {
	x := LeftRotateString("abcXYZdef", 3)
	fmt.Println(x)
}
