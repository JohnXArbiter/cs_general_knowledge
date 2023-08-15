package main

func reverseString(s []byte) {
	l, r := 0, len(s)
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
