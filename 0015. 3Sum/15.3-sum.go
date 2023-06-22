package leetcode

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// Method 1. Two Pointer (最優解)
	sort.Ints(nums)
	result, length := make([][]int, 0), len(nums)

	// 先取第一個，再使用 Two Pointer 配對看看有沒有
	for index := 1; index < length-1; index++ {
		start, end := 0, length-1
		// 如果 當前元素和前一個素相同，那就不用從頭開始
		// 因為前面的元素已經被遍歷過了，沒有必要重複再來一次
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}

		for start < index && end > index {
			// 如果和前一元素相同，直接略過，因為前面的元素已經被遍歷過了，沒有必要重複再來一次
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			// 如果和後一元素相同，直接略過，因為後面的元素已經被遍歷過了，沒有必要重複再來一次
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum := nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result

	// // Method 2. 公式
	// sort.Ints(nums)
	// return nSum(nums, 3, 0, 0)
}

func nSum(nums []int, n int, start int, target int) [][]int {
	length := len(nums)
	var res [][]int

	if n < 2 || length < n {
		return res
	}

	if n == 2 {
		small, big := start, length-1
		for small < big {
			left, right := nums[small], nums[big]
			sum := left + right

			if sum == target {
				res = append(res, []int{left, right})
				for small < big && nums[small] == left {
					small++
				}
				for small < big && nums[big] == right {
					big--
				}
			} else if sum > target {
				for small < big && nums[big] == right {
					big--
				}
			} else if sum < target {
				for small < big && nums[small] == left {
					small++
				}
			}
		}
	} else {
		i := start
		for i < length {
			now := nums[i]
			sub := nSum(nums, n-1, i+1, target-now)
			for _, s := range sub {
				s = append(s, now)
				res = append(res, s)
			}
			for i < length && nums[i] == now {
				i++
			}
		}
	}

	return res
}

// @lc code=end
