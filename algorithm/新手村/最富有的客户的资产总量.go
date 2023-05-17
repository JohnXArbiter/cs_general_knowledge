package main

import (
	"fmt"
)

func main() {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, row := range accounts {
		var fortune int
		for _, value := range row {
			fortune += value
		}
		if fortune > max {
			max = fortune
		}
	}
	return max
}
