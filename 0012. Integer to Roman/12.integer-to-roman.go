package leetcode

import "strings"

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	var sb strings.Builder
	convertMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	// 從大到小
	iterator := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, val := range iterator {
		for num >= val {
			sb.WriteString(convertMap[val])
			num -= val
		}
	}
	return sb.String()
}

// @lc code=end
