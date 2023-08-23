package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := ""
	for l := 0; l < n; l++ { // l为本次循环遍历的子串长度
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true // 只有一个字符
			} else if l == 1 {
				dp[i][j] = s[i] == s[j] // 当长度为2时，判断相不相等
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1] // 当长度大于等于3时，先判断两侧是否一样，然后就是里面也得一位的也得一样
			}
			if dp[i][j] && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
