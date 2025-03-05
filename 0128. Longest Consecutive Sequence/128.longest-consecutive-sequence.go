package leetcode

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	hashmap := make(map[int]bool)
	for _, v := range nums {
		hashmap[v] = false
	}

	maxCount := 0
	for k, _ := range hashmap {
		if v, _ := hashmap[k]; v {
			continue
		}

		count := 0
		for index := k; ; index++ {
			if _, ok := hashmap[index]; !ok {
				break
			}
			hashmap[index] = true
			count++
		}
		for index := k - 1; ; index-- {
			if _, ok := hashmap[index]; !ok {
				break
			}
			hashmap[index] = true
			count++
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}

// @lc code=end
