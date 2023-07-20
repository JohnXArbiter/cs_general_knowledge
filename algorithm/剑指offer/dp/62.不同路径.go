package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(1, 1))
}

// 优化
func uniquePaths(m int, n int) int {
	var dp = make([]int, n)
	for i := 0; i < m; i++ {
		dp[0] = 1
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

// 直接 dp
func uniquePaths2(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	var dp = make([][]int, 0)
	for i := 0; i < m; i++ {
		arr := make([]int, 0)
		for j := 0; j < n; j++ {
			arr = append(arr, 0)
		}
		dp = append(dp, arr)
	}

	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	fmt.Println(dp)

	return dp[m-1][n-1]
}
