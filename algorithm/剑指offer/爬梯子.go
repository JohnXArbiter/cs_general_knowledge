package main

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

func numWays2(n int) int {
	a, b := 1, 1
	for i := 1; i <= n; i++ {
		a, b = b%1000000007, (a+b)%1000000007
	}
	return a
}
