package main

import "fmt"

func main() {
	a := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	fmt.Println(getLeastNumbers(a, 8))
}

func getLeastNumbers(arr []int, k int) []int {
	for i := 1; i < len(arr); i++ {
		c := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] >= c {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = c
	}
	//res := make([]int, 0)
	//m := make(map[int]int)
	//for _, v := range arr {
	//	m[v] = 1
	//}
	//for _, v := range arr {
	//	if len(res) == k {
	//		break
	//	}
	//	c := m[v]
	//	if c == 1 {
	//		res = append(res, v)
	//		m[v]--
	//	}
	//}
	return arr[:k]
}
