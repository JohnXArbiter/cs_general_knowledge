package main

//func spiralOrder(matrix [][]int) []int {
//	var (
//		res         = make([]int, 0)
//		col         = len(matrix[0])
//		row         = len(matrix)
//		left, right = 0, col - 1
//		top, bottom = 0, row - 1
//		c, r        int
//	)
//	total := col * row
//	for i := 0; i < total; i++ {
//		res = append(res, matrix[r][c])
//		if c < right && r <= bottom {
//			c++
//			if c == right {
//				right--
//			}
//		} else if c == right+1 && r < bottom {
//			r++
//			if r ==  {
//
//			}
//		} else if c >= left && r == bottom {
//			c--
//		} else if {
//
//		}
//	}
//	return res
//}
