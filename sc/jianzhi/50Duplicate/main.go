package main

import "fmt"

func Duplicate1(numbers []int, duplication *[1]int) bool {
	hashset := make(map[int]int)
	for _, v := range numbers {
		if _, ok := hashset[v]; ok {
			duplication[0] = v
			return true
		} else {
			hashset[v] = 1
		}
	}
	return false

}

func Duplicate(numbers []int, duplication *[1]int) bool {
	for i := 0; i < len(numbers); i++ {
		if i == numbers[i] {
			continue
		} else {
			if numbers[numbers[i]] == numbers[i] {
				(*duplication)[0] = numbers[i]
				return true
			} else {
				swap(numbers[i], numbers[numbers[i]])
			}
		}
	}
	(*duplication)[0] = -1
	return false

}

func swap(a, b int) {
	a, b = b, a
}

func main() {
	data := []int{6, 3, 2, 0, 2, 5, 0}
	//data = []int{}
	arr := [1]int{}

	fmt.Println(Duplicate(data, &arr))
	fmt.Println(arr)
}
