package leetcode

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	// 紀錄總共跳了幾下
	jump := 0
	// 紀錄當下與之前能夠跳最遠位置
	end := 0
	// 紀錄跳最多遠
	farthest := 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		// 如果已經在能夠跳得最遠位置，則選擇往下跳，重新紀錄下次跳最遠位置
		if end == i {
			jump++
			end = farthest
		}
	}
	return jump
}

// @lc code=end
