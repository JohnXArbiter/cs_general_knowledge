package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2, 2, 3}
	fmt.Println(findLeastNumOfUniqueInts1(nums, 2))
}

// 不知道怎么改
func findLeastNumOfUniqueInts(arr []int, k int) int {
	arrMap := make(map[int]int)
	for _, value := range arr {
		arrMap[value]++
	}
	for key := range arrMap {
		if arrMap[key] == 1 && k > 0 {
			delete(arrMap, key)
			k--
		}
	}
	for k > 0 {
		min := math.MaxInt
		var keyOuter int
		for key := range arrMap {
			if arrMap[key] < min {
				keyOuter = key
				min = arrMap[key]
			}
		}
		if min == k {
			delete(arrMap, keyOuter)
			return len(arrMap)
		} else if min > k {
			temp := k
			k -= arrMap[keyOuter]
			arrMap[keyOuter] -= temp
		} else if min < k {
			k -= min
			delete(arrMap, keyOuter)
		}
	}
	return len(arrMap)
}

func findLeastNumOfUniqueInts1(arr []int, k int) int {
	var length = len(arr)
	var lenSlice = make([]int, length)
	var table = make(map[int]int, length)
	var l int
	for _, i := range arr {
		if v, exist := table[i]; exist {
			lenSlice[v]++
		} else {
			table[i] = l
			lenSlice[l] = 1
			l++
		}
	}
	if k == 0 {
		return l
	}
	sort.Ints(lenSlice)

	var idx int
	for ; k > 0; idx++ {
		if lenSlice[idx] == 0 {
			continue
		}
		k -= lenSlice[idx]
		if k < 0 {
			return l
		}
		l--
	}
	return l
}
