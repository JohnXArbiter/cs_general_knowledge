package main

import "fmt"

func main() {
	fmt.Println(maxValue([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
func maxValue(grid [][]int) int {
	var dp [][]int

	a := len(grid)
	b := len(grid[0])
	for i := 0; i <= a; i++ {
		arr := make([]int, 0)
		for j := 0; j <= b; j++ {
			arr = append(arr, 0)
		}
		dp = append(dp, arr)
	}

	for i := 1; i < a; i++ {
		dp[i][0] = dp[i-1][0] + grid[i-1][0]
	}
	for i := 1; i < b; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i-1]
	}

	for i := 1; i < a; i++ {
		for j := 1; j < b; j++ {
			dp[i][j] = Max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[a-1][b-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
