package 数组

func main() {

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	r := len(matrix)
	if r == 0 {
		return false
	}
	c := len(matrix[0])
	if c == 0 {
		return false
	}
	var i, j = r - 1, 0
	for i >= 0 && j < c {
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] {
			i--
		} else {
			return true
		}
	}
	return false
}
