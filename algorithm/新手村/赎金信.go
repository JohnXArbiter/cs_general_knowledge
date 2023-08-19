package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "bb"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	var m = [26]int32{}
	for _, s := range magazine {
		m[s-'a']++
	}
	for _, s := range ransomNote {
		m[s-'a']--
		if m[s-'a'] < 0 {
			return false
		}
	}

	return true
}
