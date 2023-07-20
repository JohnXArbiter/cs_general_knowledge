package main

import "fmt"

func main() {
	fmt.Println(rob3([]int{2, 7, 9, 3, 1}))
}

// 滚动数组优化
func rob1(nums []int) int {
	var length int
	if nums == nil {
		return 0
	}
	length = len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	a, b := nums[0], nums[1]
	for i := 2; i < length; i++ {
		a, b = b, Max2(a+nums[i], b)
	}
	return b
}

// 一维数组优化，因为能保证每次算当前的最大价值时，i-1是没偷的情况
func rob2(nums []int) int {
	var length int
	if nums == nil {
		return 0
	}
	length = len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	var dp = make([]int, length)
	dp[0], dp[1] = nums[0], Max2(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = Max2(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[length-1]
}

func rob3(nums []int) int {
	var length int
	if nums == nil {
		return 0
	}
	length = len(nums)
	if length == 0 {
		return 0
	}

	var dp = make([][]int, 0)
	for i := 0; i < 2; i++ {
		arr := make([]int, length)
		dp = append(dp, arr)
	}

	dp[0][0] = nums[0]
	for i := 1; i < length; i++ {
		dp[0][i] = dp[1][i-1] + nums[i]
		dp[1][i] = Max2(dp[0][i-1], dp[1][i-1])
	}
	return Max2(dp[0][length-1], dp[1][length-1])
}

func Max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
