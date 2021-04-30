package main

import "fmt"

func merge(data []int, s, mid, e int) int {
	i := s
	j := mid + 1
	tmp := make([]int, e-s+1)
	k := 0
	cnt := 0
	for i <= mid && j <= e {
		if data[i] <= data[j] {
			tmp[k] = data[i]

			i++
		} else {
			tmp[k] = data[j]
			cnt += mid-i+1
			j++
		}
		k++

	}
	for i <= mid {
		tmp[k] = data[i]
		i++
		k++
	}
	for j <= e {
		tmp[k] = data[j]
		j++
		k++
	}
	k = 0
	for i:=s; i<=e; i++ {
		data[i] = tmp[k]
		k++
	}
	return cnt

}

func mergeSort(data[]int, s int, e int) int {
	if e <= s {
		return 0
	}
	mid := s + (e - s) / 2
	n1 := mergeSort(data, s, mid)
	n2 := mergeSort(data, mid + 1, e)
	n3 := merge(data, s, mid, e)
	return n1 + n2 + n3

}
func InversePairs( data []int ) int {
	// write code here
	n := mergeSort(data, 0, len(data)-1)
	return n%1000000007
}


func main() {
	data := []int{1,2,3,4,5,6,7,0}
	//data = []int{4,3,2,1}
	ans := InversePairs(data)
	fmt.Println(data)
	fmt.Println(ans)
	//data := []int{1,3,2,5}
	//merge(data, 0,1,3)
	//fmt.Println(data)
}

