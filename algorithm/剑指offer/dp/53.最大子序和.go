package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray2([]int{5, 4, -1, 7, 8}))
}

func maxSubArray(nums []int) int {
	var ans = nums[0]
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

// 贪心
func maxSubArray2(nums []int) int {
	var ans = math.MinInt64
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		count += nums[i]
		if count > ans {
			ans = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return ans
}
