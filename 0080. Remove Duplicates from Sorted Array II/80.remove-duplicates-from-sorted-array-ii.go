package leetcode

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	count, prev, now, pointer := 0, -9999, -9999, 0

	for _, v := range nums {
		now = v
		if now != prev {
			count = 1
		} else {
			count++

			if count > 2 && now == prev {
				prev = now
				continue
			}
		}

		prev = now
		nums[pointer] = now
		pointer++
	}

	return pointer
}

// @lc code=end
