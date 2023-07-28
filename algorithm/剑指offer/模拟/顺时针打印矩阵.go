package main

import "fmt"

func main() {
	var matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var (
		res         = make([]int, 0)
		col         = len(matrix[0])
		row         = len(matrix)
		left, right = 0, col - 1
		top, bottom = 0, row - 1
	)
	total := col * row
	for i := 0; i < total; {
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
			i++
		}
		top++

		for j := top; j <= bottom; j++ {
			res = append(res, matrix[j][right])
			i++
		}
		right--

		for j := right; j >= left; j-- {
			res = append(res, matrix[bottom][j])
			i++
		}
		bottom--

		for j := bottom; j >= top; j-- {
			res = append(res, matrix[j][left])
			i++
		}
		left++
	}
	return res[0:total]
}
