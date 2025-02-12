package leetcode

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	size := len(nums)
	k %= size
	if k != 0 {
		doRotate(nums, 0, size-1)
		doRotate(nums, 0, k-1)
		doRotate(nums, k, size-1)
	}
}

func doRotate(nums []int, start, end int) {
	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// @lc code=end
