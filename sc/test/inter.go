package main

import "fmt"

type in interface {
	Qq()
}

type a struct {

}

func (aa *a)Qq()  {
	fmt.Println(1)
}
func (aa a)Ww()  {
	fmt.Println(1)
}
func main() {
	var aa a
	aa.Qq()

	in := &aa
	in.Qq()

	in.Ww()



}


