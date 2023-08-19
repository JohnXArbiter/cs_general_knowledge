package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

func reverseLeftWords2(s string, n int) string {
	sb := []byte(s)
	front := sb[0:n]

	sb = sb[n:]
	return string(append(sb, front...))
}

func reverseLeftWords(s string, n int) string {
	sb := []byte(s)
	reverse3(sb)

	idx := len(sb) - n
	reverse3(sb[0:idx])
	reverse3(sb[idx:])
	return string(sb)
}

func reverse3(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
