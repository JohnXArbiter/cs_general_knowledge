package main

import "strconv"

func translateNum(num int) int {
	if num < 10 {
		return 1
	}

	var numStr = strconv.Itoa(num)
	dp := make([]int, len(numStr))
	dp[0], dp[1] = 1, 1
	nums := make([]byte, 0)
	for i := range numStr {
		nums = append(nums, numStr[i])
	}

	for i := 2; i <= len(numStr); i++ {
		tmp := 10*(numStr[i-2]-'0') + (numStr[i-1] - '0')
		if tmp >= 10 && tmp < 26 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}

// 斐波那契优化
func translateNum2(num int) int {
	if num < 10 {
		return 1
	}

	var numStr = strconv.Itoa(num)
	a, b, c := 1, 1, 0
	nums := make([]byte, 0)
	for i := range numStr {
		nums = append(nums, numStr[i])
	}

	for i := 2; i <= len(numStr); i++ {
		tmp := 10*(numStr[i-2]-'0') + (numStr[i-1] - '0')
		if tmp >= 10 && tmp < 26 {
			c = a + b
		} else {
			c = b
		}
		a = b
		b = c
	}
	return c
}
