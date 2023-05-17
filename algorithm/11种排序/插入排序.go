package main

import "fmt"

func main() {
	array := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	insertionSort(array)
}

// 插入排序，就是从第二个开始，跟前面的比较，是否小于，就插到他的前面。实现的时候注意，比他大的都得往后移一位
func insertionSort(array []int) {
	length := len(array)
	for i := 1; i < length; i++ {
		ai := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if ai <= array[j] {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = ai
		fmt.Println(array)
	}
	fmt.Println(array)
}

// 插入排序优化
func insertionSortUpgraded(array []int) {
	length := len(array)
	for i := 1; i < length; i++ {
		ai := array[i]
		j := i - 1
		for array[j] <= ai {
			array[j+1] = array[j]
		}
		array[j] = ai
	}
	fmt.Println(array)
}
