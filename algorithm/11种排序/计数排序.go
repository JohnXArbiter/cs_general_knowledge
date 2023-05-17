package main

import "fmt"

// 有map的思想
func main() {
	array := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 2341, 34}
	countingSort(array)
}

func countingSort(a []int) {
	cntLen := getMax2(a) + 1
	cnt := make([]int, cntLen)

	for _, value := range a {
		cnt[value]++
	}

	index := 0
	for i := 0; i < cntLen; i++ {
		for cnt[i] > 0 {
			cnt[i]--
			a[index] = i
			index++
		}
	}
	fmt.Println(a)
}

func getMax2(a []int) int {
	var max int
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
