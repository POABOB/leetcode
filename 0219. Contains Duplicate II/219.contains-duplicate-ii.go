package leetcode

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	hashmap := make(map[int]int)
	for k1, v1 := range nums {
		if k2, ok := hashmap[v1]; ok && k1-k2 <= k {
			return true
		}
		hashmap[v1] = k1
	}
	return false
}

// @lc code=end
