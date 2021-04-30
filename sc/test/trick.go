package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.AfterFunc(8*time.Second, func() {
		fmt.Println("Golang来啦")
	})
	for {
		select {
		case <-t.C:
			fmt.Println("seekload")
			return
		}
	}
}
