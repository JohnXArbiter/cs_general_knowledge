package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
}

func add(x, y int) *int {
	z := x + y
	return &z
}
