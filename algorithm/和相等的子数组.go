package main

import "fmt"

func main() {
	array := []int{0, 0}
	t := findSubarrays1(array)
	fmt.Printf("%v", t)
}

// 参考答案
func findSubarrays1(nums []int) bool {
	mp := map[int]bool{}

	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i+1] + nums[i]
		if mp[sum] {
			return true
		}
		mp[sum] = true
	}
	return false
}
