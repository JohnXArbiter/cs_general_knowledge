package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	length := len(prices)
	dp := make([][2]int, length)
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i])))
	}
	result := int(math.Max(float64(dp[length-1][0]), float64(dp[length-1][1])))
	return result
}
