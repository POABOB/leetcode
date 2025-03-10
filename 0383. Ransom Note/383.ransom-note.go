package leetcode

/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	hashmap := make(map[byte]int, 26)
	count := 0
	for k, _ := range ransomNote {
		hashmap[ransomNote[k]]++
		count++
	}

	for k, _ := range magazine {
		if num, ok := hashmap[magazine[k]]; ok && num > 0 {
			hashmap[magazine[k]]--
			count--
		}
		if count == 0 {
			return true
		}
	}
	return false
}

// @lc code=end
