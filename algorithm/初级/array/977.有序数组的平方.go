package main

func sortedSquares(nums []int) []int {
	var length = len(nums)
	ans := make([]int, length)
	tail := length - 1
	l, r := 0, tail
	for l <= r {
		left := nums[l] * nums[l]
		right := nums[r] * nums[r]
		if left > right {
			ans[tail] = left
			l++
		} else {
			ans[tail] = right
			r--
		}
		tail--
	}
	return ans
}
