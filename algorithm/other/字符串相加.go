package main

import "fmt"

func add1(arr1 []int, arr2 []int) []int {
	l1 := len(arr1)
	l2 := len(arr2)
	maxLen := max(l1, l2) + 1
	res := make([]int, maxLen)

	carry := 0
	for i := 0; i < maxLen; i++ {
		num1 := getNum1(arr1, l1-1-i)
		num2 := getNum1(arr2, l2-1-i)

		sum := num1 + num2 + carry

		res[maxLen-1-i] = sum % 10
		carry = sum / 10
	}

	if res[0] == 0 {
		return res[1:]
	}

	return res
}

func getNum1(arr []int, index int) int {
	if index >= 0 && index < len(arr) {
		return arr[index]
	}
	return 0
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(add([]int{3, 3, 2}, []int{4, 2, 1}))
}
