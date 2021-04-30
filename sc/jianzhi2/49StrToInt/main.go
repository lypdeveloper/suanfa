package main

import (
	"fmt"
)
func StrToInt( str string ) int {
	// write code here
	if len(str) == 0 {
		return 0
	}
	ans := 0
	fuhao := 1
	has := 0
	if str[0] == '-' {
		fuhao = -1
		has = 1
	} else if str[0] == '+'{
		fuhao = 1
		has = 1
	} else if str[0]>'0' && str[0] <'9'{
		fuhao = 1
	} else {
		return 0
	}
	for i:= has; i<len(str); i++ {
		if str[i]>'0' && str[i] <'9' {
			ans = ans * 10 + int(str[i]-48)
		} else {
			return 0
		}

	}
	if fuhao == -1 {
		return -ans
	}
	return ans
}

func main() {
	fmt.Println(StrToInt(""))
}

