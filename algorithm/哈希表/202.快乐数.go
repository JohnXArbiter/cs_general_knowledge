package main

// 把每次算出来的数放到哈希表，有1就为快乐数，有重复不就是1的时候返回false
func isHappy(n int) bool {
	var m = make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
