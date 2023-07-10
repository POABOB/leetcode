package leetcode

// import "sort"

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {

	// Method 1. HashMap 計數，最基本的寫法
	// nums_len, ans := len(nums), 0
	// count := make(map[int]int, nums_len)
	// for _, v := range nums {
	// 	count[v]++
	// }

	// for _, v := range nums {
	// 	if count[v] > nums_len/2 && count[v] > ans {
	// 		ans = v
	// 	}
	// }
	// return ans

	// Method 2. Boyer–Moore majority vote algorithm(摩爾投票算法)
	target, count, nums_len := 0, 0, len(nums)
	for i := 0; i < nums_len; i++ {
		if count == 0 {
			// 當 count== 0，假設 nums[i] 是眾數，並出現一次
			target = nums[i]
			count = 1
		} else if nums[i] == target {
			// 遇到一樣就累加
			count++
		} else {
			// 沒有就遞減
			count--
		}
	}
	// 通過 ++/--，此時的count會大於0，返回的 target 一定會是目標眾數
	return target

	// Method 3. 先排序，取最中間的元素就是答案，因為出現次數最多，必然在排序後會在中間
	// sort.Ints(nums)
	// return nums[len(nums)/2]
}

// @lc code=end
