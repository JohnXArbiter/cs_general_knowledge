package main

// 直接每次跳2k个就行，然后看看离尾部是不是小于k
func reverseStr(s string, k int) string {
	if len(s) == 1 {
		return s
	}
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k < length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
