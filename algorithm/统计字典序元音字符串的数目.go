package main

import "fmt"

func main() {
	fmt.Printf("%d", countVowelStrings(33))
}

func countVowelStrings(n int) int {
	dp := [5]int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		for j := 1; j < 5; j++ {
			dp[j] += dp[j-1]
		}
	}
	count := 0
	for i := 0; i < 5; i++ {
		count += dp[i]
	}
	return count
}
