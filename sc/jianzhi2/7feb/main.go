package main

import "fmt"

func Fibonacci( n int ) (int) {
	// write code here
	if n == 1 {
		return 1
	}
	feb := []int {
		0,1,
	}

	ans := 0
	for n-2>=0 {
		ans = feb[len(feb)-1] + feb[len(feb)-2]
		feb = append(feb, ans)
		n--
	}
	fmt.Println(feb)
	fmt.Println(ans)
	return ans
}


func main() {
	Fibonacci(30)
}

