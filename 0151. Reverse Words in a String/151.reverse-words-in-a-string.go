package leetcode

import "strings"

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {

	// Method 1. 把字串根據空格切割成多筆至陣列中，然後再陸續 append 到新的陣列，最後重新組裝成字串
	// words := strings.Fields(s)
	// var ans []string
	// for i := len(words) - 1; i >= 0; i-- {
	// 	if words[i] != "" {
	// 		ans = append(ans, words[i])
	// 	}
	// }
	// return strings.Join(ans, " ")

	// Method 2. 把字串根據空格切割成多筆至陣列中，將該陣列元素 reverse，最後重新組裝成字串
	// words := strings.Fields(s)
	// i, j := 0, len(words) - 1
	// for j > i {
	//     words[i], words[j] = words[j], words[i]
	//     i++
	//     j--
	// }
	// return strings.Join(words, " ")

	// Method 3. 使用 Tow Pointer 從尾部慢慢組裝，配合 String Builder 節省記憶體空間
	var result strings.Builder
	p1, p2 := len(s), len(s)
	for p1 >= 0 {
		p1--
		if p1 >= 0 && s[p1] != ' ' {
			continue
		}
		// 遇到兩個空格，像是"a  a"
		if p1+1 == p2 {
			p2 = p1
			continue
		}
		if len(result.String()) > 0 {
			result.WriteString(" ")
		}

		result.WriteString(s[p1+1 : p2])
		p2 = p1
	}
	return result.String()
}

// @lc code=end
