package leetcode

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sIndex, tIndex := 0, 0
	for len(t) > tIndex {
		if t[tIndex] == s[sIndex] {
			sIndex++
		}
		tIndex++

		if sIndex == len(s) {
			return true
		}
	}
	return false
}

// @lc code=end
