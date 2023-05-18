package main

import "fmt"

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为0(log n)的算法。

func main() {
	var a = []int{1, 4, 6, 7, 8, 324, 764, 6456, 234324}
	fmt.Println(SearchInsert(a, 324))
}

func SearchInsert(a []int, target int) int {
	l, r := 0, len(a)-1
	ans := len(a)
	for l <= r {
		mid := (l + r) / 2
		if a[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
