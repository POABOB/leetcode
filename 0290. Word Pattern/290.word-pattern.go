package leetcode

import "strings"

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	// 紀錄 pattern 對應 word
	patternToWord := make(map[byte]string)
	// Set 紀錄被配過的組合
	wordSet := make(map[string]bool)

	for k, _ := range pattern {
		c := pattern[k]
		w := words[k]

		// 如果不是被收錄在 hasmap
		if _, ok := patternToWord[c]; !ok {
			// 又也被使用過
			if _, ok := wordSet[w]; ok {
				return false
			}
			patternToWord[c] = w
		} else {
			// hasmap 的 word 與原本 word 不同
			if patternToWord[c] != w {
				return false
			}
		}
		wordSet[w] = true
	}
	return true
}

// @lc code=end
