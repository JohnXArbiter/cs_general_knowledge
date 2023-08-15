package main

func reverseWords(s string) string {
	sb := []byte(s)
	reverse2(sb)
	slow, fast := 0, 0

	// 去掉头部空格
	for sb[fast] == ' ' {
		fast++
	}

	// 去掉中间空格
	for ; fast < len(sb); fast++ {
		if fast > 1 && sb[fast] == ' ' && sb[fast-1] == ' ' {
			continue
		}
		sb[slow] = sb[fast]
		slow++
	}

	// 去掉尾部空格
	if slow-1 > 0 && sb[slow-1] == ' ' {
		sb = sb[:slow-1]
	} else {
		sb = sb[:slow]
	}

	var idx int
	for i := 0; i < len(sb); i++ {
		if sb[i] == ' ' {
			reverse2(sb[idx:i])
			idx = i + 1
		}
		if i == len(sb)-1 {
			reverse2(sb[idx:])
		}
	}

	// 或者这样
	//i := 0
	//for i < len(b) {
	//	j := i
	//	for ; j < len(b) && b[j] != ' '; j++ {
	//	}
	//	reverse(&b, i, j-1)
	//	i = j
	//	i++
	//}

	return string(sb)
}

func reverse2(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
