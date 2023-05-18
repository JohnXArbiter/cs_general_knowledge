package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

func reverseLeftWords(s string, n int) string {
	sb := []byte(s)
	front := sb[0:n]

	sb = sb[n:]
	return string(append(sb, front...))
}
