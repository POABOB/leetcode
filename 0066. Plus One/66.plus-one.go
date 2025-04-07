package leetcode

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1

	for i := lastIndex; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}

// @lc code=end
