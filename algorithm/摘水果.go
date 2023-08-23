package main

import "fmt"

func main() {
	fmt.Println(totalFruit2([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4, 7, 8, 9, 12, 123, 132, 123123, 12312312}))
}

func totalFruit(fruits []int) int {
	var res int
	cntMap := map[int]int{}
	left := 0
	for right, x := range fruits {
		cntMap[x]++
		for len(cntMap) > 2 {
			cnt := fruits[left]
			cntMap[cnt]--
			if cntMap[cnt] == 0 {
				delete(cntMap, cnt)
			}
			left++
		}
		res = max2(res, right-left+1)
	}
	return res
}

func max2(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func totalFruit2(fruits []int) int {

	cnt := map[int]int{}
	j := 0
	for _, x := range fruits {
		cnt[x]++
		if len(cnt) > 2 {
			y := fruits[j]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			j++
		}
	}
	return len(fruits) - j
}

// 暴力
func totalFruit3(fruits []int) int {
	var res int
	length := len(fruits)

	for i, f1 := range fruits {
		count := 1
		kind := 1
		f2 := f1
		for j := i + 1; j < length; j++ {
			if fruits[j] != f1 && fruits[j] != f2 {
				f2 = fruits[j]
				kind++
			}
			if kind > 2 {
				break
			}
			count++
		}
		if count > res {
			res = count
		}
	}
	return res
}
