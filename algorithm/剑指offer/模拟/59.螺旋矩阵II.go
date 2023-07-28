package main

func generateMatrix(n int) [][]int {
	var ans [][]int
	top, bottom := 0, n-1
	left, right := 0, n-1
	loop := n * n
	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, n))
	}
	cnt := 1
	for cnt <= loop {
		for i := left; i <= right; i++ {
			ans[top][i] = cnt
			cnt++
		}
		top++

		for i := top; i <= bottom; i++ {
			ans[i][right] = cnt
			cnt++
		}
		right--

		for i := right; i >= left; i-- {
			ans[bottom][i] = cnt
			cnt++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			ans[i][left] = cnt
			cnt++
		}
		left++
	}
	return ans
}
