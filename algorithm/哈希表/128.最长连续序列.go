package main

func longestConsecutive(nums []int) int {
	var ans, cnt int
	m := map[int]bool{}
	for _, n := range nums {
		m[n] = true
	}

	for n := range m {
		if !m[n-1] { // 判断有没有比当前数字小1的数，没有才开始统计当前，否则统计了也是白统计
			i := 0
			for m[n+i] {
				cnt++
				i++
			}
			if cnt > ans {
				ans = cnt
			}
			cnt = 0
		}
	}
	return ans
}
