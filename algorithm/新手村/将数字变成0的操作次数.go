package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
}

func numberOfSteps(num int) int {
	var count int
	var temp = num
	for temp > 0 {
		if temp%2 == 0 {
			temp /= 2
			count++
		} else {
			temp -= 1
			count++
		}
	}
	return count
}
