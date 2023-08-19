package main

func main() {

}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if i2, ok := m[target-num]; ok {
			return []int{i, i2}
		} else {
			m[num] = i
		}

	}
	return []int{}
}
