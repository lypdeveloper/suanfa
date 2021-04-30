package main

import "fmt"

func mergeSort(arr *[]int, l int, r int) {
	if l >= r {//不可再分隔,只有一个元素
		return
	}
	mid := (l+r)/2
	mergeSort(arr,l,mid)
	mergeSort(arr,mid+1,r)
	if (*arr)[mid] > (*arr)[mid+1] {
		merge(arr, l, mid, r)
	}
}

//将arr[l...mid]和arr[mid+1...r]两部分进行归并
func merge(arr *[]int, l int, mid int, r int)  {
	//先把arr中[l..r]区间的值copy一份到arr2
	//注意:****这里可优化，copyarr可改为长度为r-l+1的数组,下面的赋值等操作按偏移量l来修改即可****，
	copyarr := make([]int, r+1)
	for index := l; index <= r; index++ {
		copyarr[index] = (*arr)[index]
	}
	//定义要合并的两个子数组各自目前数组内还没被合并的首位数字下标为i,j
	//初始化i，j
	i := l
	j := mid+1
	//遍历并逐个确定数组[l,r]区间内数字的顺序
	for k := l; k <= r; k++ {
		//防止i/j"越界"，应该先判断i和j的下变是否符合条件（因为两个子数组应该符合i<=mid j<=r）
		if i > mid {
			(*arr)[k] = copyarr[j]
			j++
		}else if j > r {
			(*arr)[k] = copyarr[i]
			i++
		}else if copyarr[i] < copyarr[j] {
			(*arr)[k] = copyarr[i]
			i++
		}else{
			(*arr)[k] = copyarr[j]
			j++
		}
	}
}

func main() {
	arr := []int{8,6,2,3,1,5,7,4}
	mergeSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}