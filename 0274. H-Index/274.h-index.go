package leetcode

import "sort"

/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	// 定義
	// h papers each have at least h citations
	// And n - h papers have less than h citations
	// E.g.
	// Origin: 3 0 6 1 5
	// Sorted: 0 1 3 5 6
	// 1 篇 paper 至少有 1 次 ，citations[5 - 1] = 6
	// 2 篇 paper 至少有 2 次 ，citations[5 - 2] = 5
	// 3 篇 paper 至少有 3 次 ，citations[5 - 3] = 3
	// 4 篇 paper 至少有 4 次 ，citations[5 - 4] = 1
	// 5 篇 paper 至少有 5 次 ，citations[5 - 5] = 0
	all := len(citations)
	sort.Ints(citations)
	hIndex := 0
	for i := 1; i <= all; i++ {
		if citations[all-i] >= i {
			hIndex = i
		} else {
			break
		}
	}
	return hIndex
}

// @lc code=end
