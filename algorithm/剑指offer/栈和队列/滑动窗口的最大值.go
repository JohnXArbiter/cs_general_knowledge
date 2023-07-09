package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
	// 规律：举例窗口[1,3,-1]中，其实只需要记录[3,-1]即可，因为左侧的1又小又靠左，不可能被选
	// 如此，维护一个单调窗口队列，右侧有新值加入时，将左侧小于等于它的全部舍弃，这样能保证最左侧总是当前区间的最大值，且队列严格单调递减

	res := []int{}
	queue := []int{} //存的是下标

	for i, num := range nums {
		//清理队尾
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= num {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i < k-1 {
			continue
		}
		// 保持窗口宽度
		for i-queue[0] >= k {
			queue = queue[1:]
		}
		//fmt.Println(queue)
		res = append(res, nums[queue[0]])
	}
	return res
}

// 不行，笨方法
//func maxSlidingWindow(nums []int, k int) []int {
//	var (
//		length = len(nums)
//		res    = make([]int, 0)
//	)
//	if nums == nil || length == 0 {
//		return nil
//	}
//	max := nums[0]
//	if length < k {
//		for _, num := range nums {
//			if max < num {
//				max = num
//			}
//		}
//		res = append(res, max)
//		return res
//	}
//	var slow int
//	for fast := k - 1; fast < length; fast++ {
//		max = nums[slow]
//		for i := slow; i <= fast; i++ {
//			if nums[i] >= max {
//				max = nums[i]
//			}
//		}
//		res = append(res, max)
//		slow++
//	}
//	return res
//}
