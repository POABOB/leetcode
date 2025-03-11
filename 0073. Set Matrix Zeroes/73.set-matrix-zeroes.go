package leetcode

/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	// 先記錄 0 的位置
	y, x := len(matrix), len(matrix[0])
	zeros := make([][]int, 0)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}

	// 使用 dfs 將有 0 的位置的上下左右全部設定為 0
	for _, zeroPosition := range zeros {
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{0, 1})  // 向右
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{0, -1}) // 向左
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{1, 0})  // 向下
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{-1, 0}) // 向上
	}
}

func dfs(x, y int, matrix [][]int, direction []int) {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
		return
	}
	matrix[x][y] = 0
	dfs(x+direction[0], y+direction[1], matrix, direction)
}

// @lc code=end
