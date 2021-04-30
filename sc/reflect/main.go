package main

import (
	"fmt"
	"reflect"
)

func reflectType (x interface{})  {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", t)
}

func reflectValue(x interface{}) {
	v:= reflect.ValueOf(x)
	fmt.Printf("%v \n", v)
	k:= v.Kind()
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("%v \n", ret)
	case reflect.Int64:
		ret := float32(v.Int())
		fmt.Printf("%v \n", ret)
	}
}
func main() {

	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
	sl := make([]int, 0)
	reflectType(sl)

	reflectValue(b)
}




