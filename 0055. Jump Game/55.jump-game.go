package leetcode

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	jump := 0
	for i := 0; i < n-1; i++ {
		jump = max(jump, i+nums[i])
		// 遇到卡住就直接 return，避免錯誤答案 E.g. [0, 2, 3]
		if i >= jump {
			return false
		}
	}
	return jump >= n-1
}

// @lc code=end
