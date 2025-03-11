package leetcode

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	total := (right + 1) * (down + 1)
	ans := make([]int, 0, len(matrix)*len(matrix[0]))

	for total > len(ans) {
		ans = walk(matrix, ans, up, down, left, right, 0) // 向右
		up++
		ans = walk(matrix, ans, up, down, left, right, 1) // 向下
		right--
		ans = walk(matrix, ans, up, down, left, right, 2) // 向左
		down--
		ans = walk(matrix, ans, up, down, left, right, 3) // 向上
		left++
	}
	return ans
}

func walk(matrix [][]int, ans []int, up, down, left, right, direction int) []int {
	if len(ans) == len(matrix)*len(matrix[0]) {
		return ans
	}
	if direction == 0 {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
	} else if direction == 1 {
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
	} else if direction == 2 {
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
	} else if direction == 3 {
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
	}
	return ans
}

// @lc code=end
