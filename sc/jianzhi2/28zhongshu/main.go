package main

func MoreThanHalfNum_Solution( numbers []int ) int {
	// write code here
	hash := make(map[int]int)
	for _, v := range numbers {
		hash[v]++
	}
	len := len(numbers)
	for k, v := range hash {
		if v > len /2 {
			return k
		}
	}
	return 0
}

func main() {

}

