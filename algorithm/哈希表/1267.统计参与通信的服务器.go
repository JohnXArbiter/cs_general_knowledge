package main

/*
使用两个哈希表分别记录当前列当前行的服务器数字，
所以，当我们再次去遍历网格时，
符合条件的一定是当前是1（有服务器）再判断当前行列的任意一个是不是有两台服务器，因为一台肯定不行
*/
func countServers(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	rowMap, colMap := make(map[int]int), make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowMap[i] += 1
				colMap[j] += 1
			}
		}
	}
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowMap[i] > 1 || colMap[j] > 1) {
				ans++
			}
		}
	}
	return ans
}
