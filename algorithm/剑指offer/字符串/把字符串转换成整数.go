package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	str := "4193 with words"
	fmt.Println(strToInt(str))
}

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return sign * result
}

func myAtoi(s string) (ans int) {
	state := "start"
	table := map[string][]string{
		"start":  []string{"start", "signed", "number", "end"},
		"signed": []string{"end", "end", "number", "end"},
		"number": []string{"end", "end", "number", "end"},
		"end":    []string{"end", "end", "end", "end"},
	}
	sign := 1   // 符号部分
	number := 0 // 数字部分
	getIdx := func(c byte) int {
		switch c {
		case ' ':
			return 0 // start为table[0]
		case '-', '+':
			return 1 // signed为table[1]
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return 2 // number为table[2]
		default:
			return 3 // end为table[3]
		}
		return 3
	}

	for i := range s {
		state = table[state][getIdx(s[i])] // to_state = table[from_state][new_char_idx]
		if state == "number" {
			number = number*10 + (int(s[i]) - '0')
			if sign == 1 {
				number = min(number, math.MaxInt32)
			} else if sign == -1 { // 因为math.MaxInt32和-math.MinInt32并不相等（相差1）
				number = min(number, -math.MinInt32)
			}
		} else if state == "signed" {
			if s[i] == '-' {
				sign = -1
			}
		}
	}
	return sign * number
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
