package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	base := strs[0]
	for i := 0; i < len(base); i++ {
		for _, str := range strs[1:] {
			if i >= len(str) || base[i] != str[i] {
				return base[:i]
			}
		}
	}
	return base
}

// 又是暴力😭
func longestCommonPrefix2(strs []string) string {
	if strs == nil || len(strs) == 0 || strs[0] == "" {
		return ""
	}
	length := len(strs)

	minLen := len(strs[0])
	for i := 1; i < length; i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	var ans string

	for idx := 0; idx < minLen; idx++ {
		c := []byte(strs[0])[idx]
		flag := true
		for i := 1; i < length; i++ {
			if []byte(strs[i])[idx] != c {
				flag = false
				break
			}
		}
		if !flag {
			if idx == 0 {
				return ""
			}
			return ans
		} else {
			ans += string(c)
		}
	}
	return ans
}
