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
