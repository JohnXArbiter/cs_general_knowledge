package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("AB"))
}

func titleToNumber(columnTitle string) int {
	result := 0
	gen := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += (int(columnTitle[i]) - 64) * gen
		gen *= 26
	}
	return result
}
