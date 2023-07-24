package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := []int{0}
	for i := 1; i <= amount; i++ {
		dp = append(dp, amount+1)
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = Min2(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func Min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
