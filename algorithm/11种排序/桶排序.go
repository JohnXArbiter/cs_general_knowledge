package main

import "fmt"

func main() {
	a := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	bucketSort(a)
}

func selectionSort2(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[j] <= a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

func bucketSort(a []int) {
	max := getMax(a)
	min := getMin(a)
	bucketCount := 3
	bucket := [][]int{{}, {}, {}}
	perBucket := (max - min + bucketCount) / bucketCount
	for _, value := range a {
		bucketIndex := (value - min) / perBucket
		bucket[bucketIndex] = append(bucket[bucketIndex], value)
	}
	for i := 0; i < bucketCount; i++ {
		selectionSort2(bucket[i])
	}
	index := 0
	for i := 0; i < bucketCount; i++ {
		for _, v := range bucket[i] {
			a[index] = v
			index++
		}
	}
	fmt.Println(a)
}

func getMax(a []int) int {
	var max int
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func getMin(a []int) int {
	var min int
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}
