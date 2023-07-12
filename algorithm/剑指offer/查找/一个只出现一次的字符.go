package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
}

// 用 26 长度的数组代表每个字母的个数
// 然后遍历 s 去将字母出现次数++
// 然后遍历每一位的个数，只是这里用我们的 res
func firstUniqChar(s string) byte {
	var res [26]int
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if res[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

// 开销有点大
func firstUniqChar2(s string) byte {
	var cntMap = make(map[byte]int)
	b := []byte(s)
	for _, letter := range b {
		cntMap[letter]++
	}
	for _, letter := range b {
		if cntMap[letter] == 1 {
			return letter
		}
	}
	return ' '
}

//func firstUniqChar3(s string) byte {
//	var res int
//	if s == "" {
//		return ' '
//	}
//	cntMap := make(map[byte]int)
//	b := []byte(s)
//	for _, v := range b {
//		cntMap[v]++
//		if cntMap[v] > 1 && b[res] != v {
//			res++
//		}
//	}
//	return b[res]
//}
