package main

import "strings"

func strStr(haystack string, needle string) int {
	l, r := 0, len(needle)-1

	for r < len(haystack) {
		if strings.Compare(haystack[l:r+1], needle) == 0 {
			return l
		}
		l++
		r++
	}
	return -1
}
