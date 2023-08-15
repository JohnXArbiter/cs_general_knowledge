package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	var ans [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		if n1 > target && (n1 > 0 || target > 0) {
			return ans
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] { // 对nums[j]去重
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				sum := n1 + n2 + n3 + n4
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					ans = append(ans, []int{n1, n2, n3, n4})
					l++
					r--
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				}
			}
		}
	}
	return ans
}
