package leetcode

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	y := len(matrix)
	// 鏡向
	for i := 0; i < y; i++ {
		for j := i + 1; j < y; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 水平翻轉
	for y, _ := range matrix {
		reverse(matrix[y])
	}
}

func reverse(matrixY []int) {
	left, right := 0, len(matrixY)-1
	for right > left {
		matrixY[left], matrixY[right] = matrixY[right], matrixY[left]
		left++
		right--
	}
}

// @lc code=end
