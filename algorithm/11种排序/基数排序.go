package main

import "fmt"

func main() {
	array := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 2341, 34}
	radixSort(array)
}

// 先从个位数开始比较，依据个位数进行排序，然后再比较十位数。。。
func radixSort(a []int) {
	base := 1
	maxv := getMax3(a)
	for base < maxv {
		bucket := [][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
		for _, value := range a {
			idx := value / base % 10
			bucket[idx] = append(bucket[idx], value)
		}
		index := 0
		for i := 0; i < 9; i++ {
			for _, v := range bucket[i] {
				a[index] = v
				index++
			}
		}
		base *= 10
	}
	fmt.Println(a)
}

func getMax3(a []int) int {
	var max int
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
