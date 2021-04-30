package main

import "fmt"

func a() func(){
	name := "zzz"
	return func(){
		fmt.Println("hh",name)
	}
}

func main() {
	f := a()
	f()
}

