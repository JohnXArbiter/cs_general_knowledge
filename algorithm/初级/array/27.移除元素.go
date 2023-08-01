package main

import "fmt"

func main() {
	fmt.Println(removeElement2([]int{1, 2, 3, 1, 1, 1, 45, 35}, 1))
}

// 数组移动
func removeElement(nums []int, val int) int {
	var ans = len(nums)
	length := ans - 1
	for i := 0; i <= length; {
		if nums[i] == val {
			if i == length {
				return ans - 1
			}
			for j := i; j < length; j++ {
				nums[j] = nums[j+1]
			}
			ans--
			length--
		} else {
			i++
		}
	}
	return ans
}

func removeElement2(nums []int, val int) int {
	var slow int
	length := len(nums)
	for fast := 0; fast < length; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
