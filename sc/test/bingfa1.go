package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func main() {
	//chain = "main"
	//A()
	//fmt.Println(chain)
	var b interface{}
	a := "11"
	b = a
	fmt.Println(b.(float64))
}
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}
