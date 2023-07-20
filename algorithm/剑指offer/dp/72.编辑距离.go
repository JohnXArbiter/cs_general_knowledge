package main

import "fmt"

func main() {
	fmt.Println(minDistance2("horse", "ros"))
}

// 优化空间
func minDistance(word1 string, word2 string) int {
	a := len(word1)
	b := len(word2)

	dp := make([]int, b+1)
	for i := 1; i <= b; i++ {
		dp[i] = i
	}

	for i := 1; i <= a; i++ {
		tmp := dp[0]
		dp[0] = i
		for j := 1; j <= b; j++ {
			pre := tmp
			tmp = dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = pre
			} else {
				dp[j] = Min(pre, Min(dp[j-1], dp[j])) + 1
			}
		}
	}
	return dp[b]
}

// dp
func minDistance2(word1 string, word2 string) int {
	a := len(word1)
	b := len(word2)

	dp := make([][]int, 0)
	for i := 0; i <= a; i++ {
		arr := make([]int, 0)
		for j := 0; j <= b; j++ {
			arr = append(arr, 0)
		}
		dp = append(dp, arr)
	}
	for i := 1; i <= a; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for i := 1; i <= b; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j-1], Min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[a][b]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
