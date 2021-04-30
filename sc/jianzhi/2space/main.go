package main

import (
	"fmt"
)

func replaceSpace1(s string) string {
	//write your code here
	ans := ""
	for _, v := range s {
		if string(v) == " " {
			ans += "%20"
		} else {
			ans += string(v)
		}
	}
	return ans

}

func replaceSpace(s string) string {
	//write your code here
	str := []byte(s)
	bLen := len(str)
	if bLen == 0 {
		return ""
	}
	for i := 0; i < bLen; i++ {
		if string(str[i]) == " " {
			str = append(str, []byte("  ")...)
		}
	}
	aLen := len(str)
	fmt.Println(bLen, aLen)
	i := bLen - 1
	j := aLen - 1
	for i >= 0 && i <= j {
		if str[i] == byte(' ') {
			str[j] = byte('0')
			str[j-1] = byte('2')
			str[j-2] = byte('%')
			j -= 3
		} else {
			str[j] = str[i]
			j--
		}
		i--

	}
	fmt.Println(string(str))
	return string(str)

}

func main() {
	replaceSpace("qw er r")
	//fmt.Println(x)
}
