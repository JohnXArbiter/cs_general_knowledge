package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%s", fizzBuzz(6))
}

func fizzBuzz(n int) []string {
	var arr []string
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")

		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return arr
}
