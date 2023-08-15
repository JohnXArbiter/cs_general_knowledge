package main

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-1; i++ {
		n1 := nums[i]
		if n1 > 0 {
			return ans
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, length
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 > 0 {
				r--
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				ans = append(ans, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			}
		}

	}
	return ans
}
