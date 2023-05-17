package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 5, 1, 3, 123, 8, 234, 3, 56, 1, 123, 312, 452, 34}
	shellSort(a)
}

// 相当于优化过的插入排序，我们用 gap=长度/2 规定为一共有几组
// 那么一开始，肯定有 gap 组，并且每组是2个（因为只除过一次2），这时间隔就是gap个
// 第一次，每组两两排序
// 第二次，gap=长度/2/2，那么组数变少，一组为4个，间隔为gap个，以此类推
func shellSort(a []int) {
	length := len(a)
	gap := length / 2
	for gap > 0 {
		for i := gap; i < length; i++ {
			x := a[i]
			j := i
			for j >= gap {
				if x < a[j-gap] {
					a[j] = a[j-gap]
				} else {
					break
				}
				j -= gap
			}
			a[j] = x
		}
		gap /= 2
	}
	fmt.Println(a)
}
