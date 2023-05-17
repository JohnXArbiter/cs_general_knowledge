package main

import "fmt"

func main() {
	a := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	selectionSort(a)
}

// 选择排序，就是每次遍历数组，每次找到最小/最大的，将其放到前面。
// 外层循环n-1次，内层循环n次（每次从外层+1次开始）
func selectionSort(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
		fmt.Println(a)
	}
	fmt.Println(a)
}
