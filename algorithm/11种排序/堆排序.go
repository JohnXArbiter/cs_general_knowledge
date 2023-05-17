package main

import "fmt"

func main() {
	array := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	HeapSort(array)
}

func maxHeapify(heap []int, start, end int) {
	son := start * 2
	for son <= end {
		if son+1 <= end && heap[son+1] > heap[son] {
			son++
		}
		if heap[son] > heap[start] {
			heap[start], heap[son] = heap[son], heap[start]
			start, son = son, son*2
		} else {
			break
		}
	}
}

func HeapSort(a []int) {
	heap := []int{0} // 使用哨兵位置
	heap = append(heap, a...)
	root := 1
	l := len(heap)
	for i := l / 2; i >= root; i-- {
		maxHeapify(heap, i, l-1)
	}
	for i := l - 1; i >= root; i-- {
		heap[i], heap[root] = heap[root], heap[i]
		maxHeapify(heap, root, i-1)
	}
	fmt.Println(a)
}
