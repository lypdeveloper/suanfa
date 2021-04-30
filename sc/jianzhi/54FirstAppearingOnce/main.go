package main

import (
	"fmt"
	"strings"
)
func FirstAppearingOnce( str string ) byte {
	// write code here
	once := make(map[int32]int)
	for k, v := range str {

		if _, ok := once[v]; !ok && strings.LastIndex(str, string(v)) == k {
			return byte(v)
		}
		once[v] = 1
	}
	return byte('#')

}


func main() {
	ch1 := make(chan int,1)
	ch1<-10
	x := <-ch1
	fmt.Println(x)
}

