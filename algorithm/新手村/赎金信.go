package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "bb"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	list := [26]int{}
	for _, value := range magazine {
		list[value-'a']++
	}
	for _, value := range ransomNote {
		list[value-'a']--
		if list[value-'a'] < 0 {
			return false
		}
	}
	return true
}
