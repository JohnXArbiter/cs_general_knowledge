package main

import "fmt"

func main() {
	array := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	bubbleSort(array)
}

// 冒泡排序，外层循环n-1次，每次内层循环n-1-外层次，并从头遍历到尾比较相邻大小，左大换右
func bubbleSort(a []int) {
	length := len(a)
	for i := length - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
		fmt.Println(a)
	}
	fmt.Println(a)
}
