package leetcode

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = nums[i-1] * arr[i-1]
	}

	right := 1
	for i := n - 2; i >= 0; i-- {
		right = right * nums[i+1]
		arr[i] *= right
	}
	return arr
}

// @lc code=end
