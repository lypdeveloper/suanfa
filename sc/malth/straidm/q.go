package main

import (
	"fmt"
)

func solution(s string, a string) (ans int) {

	table := make(map[uint8]int, 0)
	invalid := 0
	for j:=0; j< len(a); j++ {
		table[a[j]] ++
	}
	i := 0
	for ; i< len(a); i++ {
		table[s[i]]--
		if table[s[i]] < 0 {
			invalid ++
		}
	}
	for ; i<len(s); i++ {
		if invalid == 0 {
			return i-len(a)
		}
		table[s[i]]--
		if table[s[i]] < 0 {
			invalid ++
		}

		if table[s[i-len(a)]] < 0 {
			invalid --
		}
		table[s[i-len(a)]]++
	}
	if invalid==0 {
		return i - len(a)
	}else {
		return -1
	}

}

func main() {
	s := "aaaqwert"
	a := "we"
	ans := solution(s, a)
	fmt.Println(ans)


}