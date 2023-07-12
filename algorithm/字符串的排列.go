package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("hello", "hooholleoooleh"))
}

func checkInclusion(s1 string, s2 string) bool {
	var (
		win   = make(map[byte]int)
		s1Cnt = make(map[byte]int)
		s1len = len(s1)
		cnt   = 0
		left  = -1
	)

	for i := range s1 {
		s1Cnt[s1[i]]++
	}

	for right := range s2 {
		win[s2[right]]++
		if win[s2[right]] == s1Cnt[s2[right]] {
			cnt++
		}

		if right-left < s1len {
			continue
		}

		if cnt == len(s1Cnt) {
			return true
		}

		left++
		if win[s2[left]] == s1Cnt[s2[left]] {
			cnt--
		}
		win[s2[left]]--
	}
	return false
}
