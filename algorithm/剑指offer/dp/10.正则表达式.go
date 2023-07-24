package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a."))
}

func isMatch(s string, p string) bool {
	if s == "" || p == "" {
		return false
	}
	if p[0] == '*' {
		return true
	}
	slen := len(s)
	plen := len(p)
	dp := make([][]bool, slen+1)
	for i := range dp {
		dp[i] = make([]bool, plen+1)
	}
	dp[0][0] = true
	for i := 1; i <= plen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			if s[i-1] == p[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if j != 1 && p[j-2] != '.' && p[j-2] != s[i-1] {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]
				}
			}
		}
	}
	return dp[slen][plen]
}
