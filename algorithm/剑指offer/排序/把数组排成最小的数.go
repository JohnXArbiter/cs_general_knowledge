package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minNumber([]int{3, 30, 34, 1, 52, 25, 4, 6, 21}))
}

func minNumber(nums []int) string {
	quickSort(nums, 0, len(nums)-1)
	res := ""
	for i := 0; i < len(nums); i++ {
		res += fmt.Sprintf("%d", nums[i])
	}
	return res
}
func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	p := nums[start]
	i, j := start, end
	for i < j {
		for i < j && compair(nums[j], p) {
			j--
		}
		nums[i] = nums[j]
		for i < j && compair(p, nums[i]) {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = p
	quickSort(nums, start, i-1)
	quickSort(nums, i+1, end)
}
func compair(a, b int) bool {
	stra, strb := strconv.Itoa(a), strconv.Itoa(b)
	if stra+strb >= strb+stra {
		return true
	} else {
		return false
	}
}
