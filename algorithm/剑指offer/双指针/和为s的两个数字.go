package main

func main() {

}

// 利用二分去搜索另一个数
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		isa, err := binarySearch(nums, a)
		if err != nil {
			return nil
		}
		if isa == a {
			return []int{nums[i], a}
		}
	}
	return nil
}

func binarySearch(nums []int, target int) (int, error) {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			return nums[mid], nil
		}
	}
	return 0, nil
}

func twoSum2(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		tmp := nums[i] + nums[j]
		if tmp == target {
			return []int{nums[i], nums[j]}
		} else if tmp > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	numsMap := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := numsMap[target-v]; ok {
			return []int{target - v, v}
		}
		if _, ok := numsMap[v]; !ok {
			numsMap[v] = struct{}{}
		}
	}
	return nil
}
