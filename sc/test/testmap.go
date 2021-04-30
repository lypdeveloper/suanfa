package main

import "fmt"

type node struct {
	M map[string]int
}

func main() {
	m := make(map[string]int)
	a := node{m}
	a.M["1"] = 2
	fmt.Println(a)
}
