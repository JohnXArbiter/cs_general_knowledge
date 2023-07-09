package main

import "fmt"

func main() {
	//strings := make([]string, 0)
	var res string
	res = res + "sad"
	fmt.Println(res)
}

func exchange(nums []int) []int {
	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j]%2 == 0 {
			continue
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}
