package main

import "fmt"

func main() {
	array := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	mergeSort(array, 0, len(array)-1)
}

// 这里是分治思想，将元素最少拆为两个一组，然后再慢慢合并排
func mergeSort(array []int, start, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSort(array, start, mid)
	mergeSort(array, mid+1, end)
	merge(array, start, mid, end)
	fmt.Println(array)
}

// 前一半就从前头开始，后一半从mid开始，都从两半的头开始跨组比较，因为同组的已经有序，每次把最小的放进临时数组里
func merge(array []int, start, mid, end int) {
	temp := make([]int, 0)
	l := start
	r := mid + 1
	for l <= mid && r <= end {
		if array[l] <= array[r] {
			temp = append(temp, array[l])
			l++
		} else {
			temp = append(temp, array[r])
			r++
		}
	}
	// 循环结束，至少有一边是全部已经加入临时表的，所以直接将后面的（反正也是排序好的），往后追加就行
	temp = append(temp, array[l:mid+1]...)
	temp = append(temp, array[r:end+1]...)
	// 替换
	for i := start; i < end+1; i++ {
		array[i] = temp[i-start]
	}
}
