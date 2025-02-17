package leetcode

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	m := len(strs)
	n := len(strs[0])

	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			thisStr, prevStr := strs[row], strs[row-1]
			// 如果長度比 col 小，就不用往後比較了
			// 如果 thisStr[col] != prevStr[col]，找到不同的字元，return
			if col >= len(thisStr) || col >= len(prevStr) || thisStr[col] != prevStr[col] {
				return strs[row][:col]
			}
		}
	}
	return strs[0]
}

// @lc code=end
