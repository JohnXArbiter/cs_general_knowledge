package main

import "fmt"

func main() {
	fmt.Println(isValid("){"))
}

func isValid(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}

	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, m[s[i]])
		} else if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if len(stack) > 0 && stack[len(stack)-1] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
