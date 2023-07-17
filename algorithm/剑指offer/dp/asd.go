package main

import "fmt"

func main() {
	fmt.Println(fb(5))
}

func fb(n int) int {
	var ans = []int{0, 1}
	if n < 2 {
		return ans[n]
	}
	for i := 2; i <= n; i++ {
		ans = append(ans, (ans[i-2]+ans[i-1])%1000000007)
	}
	return ans[n]
}
