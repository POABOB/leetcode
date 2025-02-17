package leetcode

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	convertMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := convertMap[s[len(s)-1]]
	for i := 0; i < len(s)-1; i++ {
		before := convertMap[s[i]]
		after := convertMap[s[i+1]]
		if after > before {
			total -= before
		} else {
			total += before
		}
	}
	return total
}

// @lc code=end
