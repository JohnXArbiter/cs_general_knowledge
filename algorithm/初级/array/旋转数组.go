package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	length := len(nums)
	var nums2 = make([]int, length)
	for index, num := range nums {
		nums2[(index+k)%length] = num
	}
	copy(nums, nums2)
}
