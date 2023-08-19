package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
func minWindow(s string, t string) string {
	var slen = len(s)
	var tlen = len(t)
	tmap := map[int32]int{}
	for _, e := range t {
		tmap[e]++
	}

	left := -1
	for i, e := range s {
		if _, ok := tmap[e]; ok {
			left = i - 1
			break
		}
	}
	var start, end int
	cntMap := map[int32]int{}
	min := len(s)
	cmax := 0
	for right, e := range s {
		if right == slen-1 && slen-left == tlen+1 {
			break
		}
		cntMap[e]++
		if cntMap[e] == tmap[e] {
			cmax++
		}
		if cmax == tlen {
			cntMap[int32(s[left+1])]--
			if right-left < min {
				min = right - left
				start, end = left, right
			}
			left++
			for {
				if _, ok := tmap[int32(s[left+1])]; ok {
					break
				}
				left++
			}
		}
	}
	return s[start+1 : end+1]
}

func minWindow2(s string, t string) string {
	window := make(map[byte]int, 0)
	need := make(map[byte]int, 0)

	left, match := -1, 0
	start, end, min := 0, 0, len(s)+1

	for i := range t {
		need[t[i]]++
	}

	for right := 0; right < len(s); right++ {
		// 1. 直接将s[right]加入到区间，形成（left, right]
		ch1 := s[right]
		window[ch1]++

		//  2. 更新状态
		if window[ch1] == need[ch1] {
			match++
		}

		// 3. 超出区间，或者满足条件
		for match == len(need) {
			if right-left < min {
				start, end = left, right
				min = right - left
			}

			// 4. 移除s[++left]，更新状态
			left++
			ch2 := s[left]
			if window[ch2] == need[ch2] {
				match--
			}
			window[ch2]--
		}
	}

	return s[start+1 : end+1]
}
