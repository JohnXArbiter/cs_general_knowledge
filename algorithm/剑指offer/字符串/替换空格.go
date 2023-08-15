package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	sb := []byte(s)
	newStr := ""
	for _, value := range sb {
		vb := string(value)
		if vb == " " {
			newStr += "%20"
		} else {
			newStr += string(value)
		}
	}
	return newStr
}

func replaceSpaceGO(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

/*
* 不使用额外空间：先看看有几个空格，然后把计数*2的空间加到原来的切片中
* 因为原来空格就是一个空间了，所以是*2
* 然后做指针指向原来的尾部，又指针就是新尾部，一同左移
* 如果碰到空格，右边就换为 %20
 */
func replaceSpace2(s string) string {
	sb := []byte(s)
	spaceCnt := 0
	for _, ss := range s {
		if ss == ' ' {
			spaceCnt++
		}
	}
	sb = append(sb, make([]byte, spaceCnt*2)...)
	l, r := len(s)-1, len(sb)-1
	for l < r {
		if sb[l] == ' ' {
			sb[r] = '0'
			sb[r-1] = '2'
			sb[r-2] = '%'
			r -= 2
		} else {
			sb[r] = sb[l]
		}
		l--
		r--
	}
	return string(sb)
}
