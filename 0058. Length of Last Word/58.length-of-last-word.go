package leetcode

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	left, right := len(s)-1, len(s)-1
	for left >= 0 {
		if s[left] == ' ' && s[right] == ' ' {
			left--
			right--
		} else if s[left] == ' ' {
			break
		} else {
			left--
		}
	}
	return right - left
}

// @lc code=end
