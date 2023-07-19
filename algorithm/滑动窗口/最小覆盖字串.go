package main

import "fmt"

func main() {
	fmt.Println(minWindow("DFGBWCEAAFGSDFBC", "ABC"))
}

func minWindow(s string, t string) string {
	window := make(map[int32]int, 0)
	need := make(map[int32]int, 0)

	left, match := -1, 0
	start, end, min := 0, 0, len(s)+1

	for _, e := range t {
		need[e]++
	}

	for right, e := range s {
		window[e]++
		if window[e] == need[e] {
			match++
		}
		for match == len(need) {
			if right-left < min {
				start, end = left, right
				min = right - left
			}
			left++
			ch2 := int32(s[left])
			if window[ch2] == need[ch2] {
				match--
			}
			window[ch2]--
		}
	}
	return s[start+1 : end+1]
}
