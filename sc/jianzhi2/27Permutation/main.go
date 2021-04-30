package main

import (
	"fmt"
	"sort"
)

func arrayUnique (arr *[]string) {
	m := make(map[string]int)
	o := make([]string, 0)
	for _, v := range *arr {
		m[v] = 1
	}
	for k, _ := range m {
		o = append(o, k)
	}
	*arr = o
}
func Permutation( str string ) []string {
	// write code here
	arrStr := []byte(str)
	ans := make([]string, 0)
	dfs(arrStr, 0, len(arrStr) -1,  &ans)
	//fmt.Println(ans)
	arrayUnique(&ans)
	sort.Strings(ans)

	return ans

}
func dfs(str []byte, left, right int, ans *[]string) {
	if left == right {
		*ans = append(*ans, string(str))
	}
	for i :=left; i<=right; i++ {
		str[left], str[i] = str[i], str[left]
		dfs(str, left+1, right, ans)
		str[left], str[i] = str[i], str[left]
	}
}

func main() {


	str := "abcdef"

	ans := Permutation(str)
	fmt.Println(ans)
}

