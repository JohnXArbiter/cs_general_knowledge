package main

func main() {

}

/*
* 使用哈希法，先遍历前两个数组把每个组合的计数记下来，
* 然后再遍历后面两个数组，找到符合 0 - (a + b) 的计数，
* 加到 count 里面
* 为什么不是直接 count++，因为第一次循环的遍历就是组合数
 */
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var m, ans = map[int]int{}, 0
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m[n1+n2]++
		}
	}

	for _, n1 := range nums3 {
		for _, n2 := range nums4 {
			ans += m[0-(n1+n2)]
		}
	}
	return ans
}
