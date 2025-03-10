package leetcode

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte, 26)
	tMap := make(map[byte]byte, 26)

	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; !ok {
			sMap[s[i]] = t[i]
		}
		if sMap[s[i]] != t[i] {
			return false
		}

		if _, ok := tMap[t[i]]; !ok {
			tMap[t[i]] = s[i]
		}
		if tMap[t[i]] != s[i] {
			return false
		}
	}
	return true
}

// @lc code=end
