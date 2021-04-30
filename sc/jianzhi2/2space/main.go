package main

import (
	"fmt"
)

//请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
func replaceSpace(s string) string {
	//write your code here
	if len(s) == 0 {
		return ""
	}
	str := []byte(s)
	i := len(str) -1
	for _,v:= range str {
		//fmt.Println(v)
		if v == ' ' {
			str = append(str, []byte("  ")...)
		}
	}
	//fmt.Println(string(str))
	j := len(str)-1
	for i>=0 {
		if str[i] == byte(' '){
			str[j] = byte('0')
			str[j-1] = byte('2')
			str[j-2] = byte('%')
			i--
			j-=3
		} else if str[i] != byte(' ') {
			str[j] = str[i]
			i--
			j--
		}
		//println(i, j)
	}
	return string(str)

}

func main() {
	x := replaceSpace("qw er r")
	fmt.Println(x)
}
