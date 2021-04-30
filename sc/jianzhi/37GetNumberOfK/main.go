package main

import (
	"fmt"
)

func _find(data []int, s, e int, k int) int {
	if e == s {
		if data[e] == k {
			return 1
		}

	}
	cnt := 0
	cc := 0
	mid := s + (e-s)/2
	if data[mid] == k {
		cnt++
		c1 := _find(data, s, mid, k)
		c2 := _find(data, mid+1, e, k)
		cc = c1+c2
	} else if data[mid] < k {
		cc = _find(data, mid+1, e, k)
	} else {
		cc = _find(data, s, mid, k)
	}
	return cc + cnt
}

func GetNumberOfK_old( data []int ,  k int ) int {
	// write code here
	s := 0
	e := len(data) - 1
	cnt := _find(data, s, e, k)
	return cnt

}

func GetNumberOfK( data []int ,  k int ) int {
	// write code here
	//左界
	s := 0
	e := len(data) - 1
	for s <= e {
		mid := s+(e-s)/2
		if data[mid] < k {
			s = mid + 1
		} else if data[mid] > k {
			e = mid - 1
		} else {
			e = mid - 1
		}
	}
	fmt.Println("left " , s)
	lb := s

	//右界
	s = 0
	e = len(data) - 1
	for s <= e {
		mid := s+(e-s)/2
		if data[mid] < k {
			s = mid + 1
		} else if data[mid] > k {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	rb := s
	fmt.Println("right " , s)

	return  rb - lb

}


func main() {
	data := []int{1,2,2, 3}
	k := GetNumberOfK(data, 2)
	fmt.Println(k)
}

