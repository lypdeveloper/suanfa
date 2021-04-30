package main

import "strings"

func reStr(strs *string, i, j int) {
	str := []byte(*strs)
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	*strs =  string(str)
}
func ReverseSentence( ReverseSentence string ) string {
	// write code here
	ans := ""
	arrStr := []byte(ReverseSentence)
	reStr(&ReverseSentence, 0, len(arrStr) -1)
	//fmt.Println(ReverseSentence)

	split := strings.Split(ReverseSentence, " ")
	for _, v := range split {

		reStr(&v, 0, len(v)-1)
		ans += v + " "
	}
	strings.TrimRight(ans, " ")
	return ans

}


func main() {
	 ReverseSentence("abc def")

	//fmt.Println(x)
}

