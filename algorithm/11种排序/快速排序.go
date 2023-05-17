package main

import (
	"fmt"
	"math/rand"
)

// 选取第一个数为基准数，然后向后遍历，将比基准数小的数都放在左边，大于基准数的放右边
// 时间复杂度为 O(nlogn)
func main() {
	a := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

// 选择当前start为基准，然后从start+1开始循环到end，用j标识第一个大于基准值的变量，i不断向后
// 这里使用随机快速排序，基准数随机选择，不然每次从第一个开始在升序排好的情况浪费性能时间复杂度为O(n^2)
func quickSortPivot(a []int, start, end int) int {
	random := rand.Intn(end-start+1) + start
	a[start], a[random] = a[random], a[start]
	pivot := start
	j := start + 1
	for i := start + 1; i < end+1; i++ {
		// 只要当前a[i]小于基准值，那就说明满足升序排列，则j肯定也向后
		// 不满足的话，j不加，就能代表当前下标位置的值是大于基准值的，在下一次找到小于基准值的就交换位置
		if a[i] <= a[pivot] {
			a[i], a[j] = a[j], a[i]
			j++
		}
	}
	// 由于找到小于基准值的时候ij是一起移动的，反之则不移动，所以j一定是第一个大于基准值的数的下标
	// 那就把基准值放到大于基准值之前，那前面都是小于基准值的了
	a[pivot], a[j-1] = a[j-1], a[pivot]
	pivot = j - 1
	//fmt.Println(a)
	return pivot
}

// 基于基准值的位置进行分治
func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}
	pivot := quickSortPivot(a, start, end)
	quickSort(a, start, pivot-1)
	quickSort(a, pivot+1, end)
}
