package main

import "fmt"

func main() {
	array := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	fmt.Println(BinarySearch3(array, 17))
}

func BinarySearch(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//要查找的数，范围应该在 leftIndex 到 middle+1
		BinarySearch(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//要查找的数，范围应该在 middle+1 到 rightIndex
		BinarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为：%v \n", middle)
	}
}

func BinarySearch2(a []int, target int) int {
	start, end := 0, len(a)-1
	for start <= end {
		m := start + (end-start)/2
		//fmt.Println(start, end, m, a[m], target)
		if a[m] == target {
			return m
		}
		if target > a[m] {
			start = m + 1
		}
		if target < a[m] {
			end = m - 1
		}
	}
	return -1
}

func BinarySearch3(a []int, target int) int {
	start, end := 0, len(a)-1
	for start <= end {
		mid := (start + end) / 2
		if target == a[mid] {
			return mid
		} else if target < a[mid] {
			end = mid - 1
		} else if target > a[mid] {
			start = mid + 1
		}
	}
	return -1
}
