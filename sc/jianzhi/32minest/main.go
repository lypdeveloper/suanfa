package main

import (
	"fmt"
	"sort"
	"strconv"
)


type sumsort []int
func (p sumsort) Len() int           { return len(p) }
func (p sumsort) Less(i, j int) bool {
	si:= strconv.Itoa(p[i])
	sj:= strconv.Itoa(p[j])
	sij, _ := strconv.Atoi(si+sj)
	sji ,_ := strconv.Atoi(sj+si)
	if sij <= sji {
		return true
	} else {
		return false
	}
}
func (p sumsort) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func PrintMinNumber( numbers []int ) string {
	// write code here
	sort.Sort(sumsort(numbers))
	ans := ""
	for _, v := range numbers {
		ans += strconv.Itoa(v)
	}

	return ans

}

func main() {
	ans := PrintMinNumber([]int{3,32,321})
	fmt.Println(ans)
}

