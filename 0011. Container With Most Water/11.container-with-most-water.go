package leetcode

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for right > left {
		hight := min(height[left], height[right])
		width := right - left
		area := hight * width
		if area > maxArea {
			maxArea = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}

// @lc code=end
