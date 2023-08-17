package main

func main() {

}

// 给定一个数组arr，表示连续n天的股价，数组下标表示第几天指标X︰任意两天的股价之和–此两天间隔的天数
// 比如
// 第3天，价格是10 第9天，价格是30
// 那么第3天和第9天的指标X= 10 + 30 - (9 - 3) = 34返回arr中最大的指标X
func maxXFromStock(a []int) int {
	if a == nil || len(a) < 2 {
		return -1
	}
	preBest := a[0]
	res := a[0]
	for i := 0; i < len(a); i++ {
		res = Max(res, preBest+a[i]-i)
		preBest = Max(preBest, a[i]+i)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 看成 a[i]，a[j]，答案变成 a[i] + i - a[j] - j
