package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen2(6, []int{10, 2, 3}))
}

func minSubArrayLen(target int, nums []int) int {
	var start, end int
	min := math.MaxInt64
	slow, cnt := 0, 0
	length := len(nums)
	for fast := 0; fast < length; fast++ {
		cnt += nums[fast]
		for cnt >= target {
			if fast-slow+1 < min {
				start, end = slow, fast
				min = fast - slow + 1
			}
			cnt -= nums[slow]
			slow++
		}
	}
	fmt.Println(end, start)
	if min == math.MaxInt64 {
		return 0
	}
	return min
}

func minSubArrayLen2(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
