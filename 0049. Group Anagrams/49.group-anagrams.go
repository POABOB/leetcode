package leetcode

import "strings"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	answer := make([][]string, 0)
	hashmap := make(map[string]int)
	for _, str := range strs {
		vEncode := encode(str)
		if index, ok := hashmap[vEncode]; ok {
			answer[index] = append(answer[index], str)
		} else {
			hashmap[vEncode] = len(answer)
			answer = append(answer, []string{str})
		}
	}
	return answer
}

func encode(str string) string {
	count := make([]byte, 26)
	for _, v := range str {
		count[v-'a']++
	}
	var builder strings.Builder
	for _, v := range count {
		builder.WriteByte(v + '0') // 轉成字元
	}
	return builder.String()
}

// @lc code=end
