package main

import "strings"

// 从尾部开始，fast指针指到一个单词的头，slow指针指到尾部
func reverseWords(s string) string {
	var res string
	i := len(s) - 1
	j := i
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res = res + s[i+1:j+1] + " "
	}
	return strings.TrimRight(res, " ")
}

// 正向保存，反向遍历重构
func reverseWords2(s string) string {
	var words = make([]string, 0)
	tail := len(s) - 1
	var i, j int
	for i <= tail {
		for i <= tail && s[i] == ' ' {
			i++
		}
		j = i
		for i <= tail && s[i] != ' ' {
			i++
		}
		words = append(words, s[j:i])
	}
	var res string
	for i := len(words) - 1; i >= 0; i-- {
		res = res + words[i] + " "
	}
	return strings.Trim(res, " ")
}
