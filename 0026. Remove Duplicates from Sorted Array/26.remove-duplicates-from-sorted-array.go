package leetcode

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	fast, slow := 0, 0
	for len(nums) > fast {
		if nums[slow] != nums[fast] {
			nums[slow] = nums[fast]
		}

		for len(nums) > fast && nums[fast] == nums[slow] {
			fast++
		}
		slow++
	}
	return slow
}

// @lc code=end
