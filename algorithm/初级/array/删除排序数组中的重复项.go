package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 6, 7}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != nums[i] {
					nums[i] = nums[j]
					break
				}
			}
		} else if nums[i] < nums[i-1] {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i-1] {
					nums[i] = nums[j]
					break
				}
			}
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		} else {
			break
		}
	}
	return count
}

// 快慢“指针”
// slow永远指向当前找到的不重复放到头位置的位置，fast当找到不同值的时候往后遍历
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
