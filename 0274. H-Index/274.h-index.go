package leetcode

import "sort"

/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	c_len := len(citations)
	sort.Ints(citations)

	for _, v := range citations {
		if c_len > v {
			c_len--
		}
	}

	return c_len
}

// @lc code=end
