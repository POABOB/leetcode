package leetcode

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	sCount := count(s)
	tCount := count(t)
	for i := 0; i < 26; i++ {
		if sCount[i] != tCount[i] {
			return false
		}
	}
	return true
}

func count(s string) []int {
	charCount := make([]int, 26)
	for _, v := range s {
		charCount[v-'a']++
	}
	return charCount
}

// @lc code=end
