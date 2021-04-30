package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("qq","ww")
	fmt.Println(os.Getenv("qq"))
}