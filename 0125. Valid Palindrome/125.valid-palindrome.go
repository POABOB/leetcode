package leetcode

import (
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for right > left {
		if !isLetterOrDigital(rune(s[left])) {
			left++
			continue
		} else if !isLetterOrDigital(rune(s[right])) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetterOrDigital(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsDigit(char)
}

// @lc code=end
