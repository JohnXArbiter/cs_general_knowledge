package main

func minCostClimbingStairs(cost []int) (ans int) {
	cost = append(cost, 0)
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return cost[len(cost)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
