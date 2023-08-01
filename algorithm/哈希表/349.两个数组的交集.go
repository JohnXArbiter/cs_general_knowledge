package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{2, 1}, []int{2, 3}))
}
func intersection(nums1 []int, nums2 []int) []int {
	var record = make(map[int]struct{})
	for _, v := range nums1 {
		record[v] = struct{}{}
	}
	ans := make([]int, 0)
	for _, k := range nums2 {
		if _, ok := record[k]; ok {
			ans = append(ans, k)
			delete(record, k)
		}
	}
	return ans
}
