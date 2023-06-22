package leetcode

import "sort"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// /** Method 1.
	//  * 透過 HashMap 儲存 nums 的個別元素 (map[元素]index)
	//  * 再透過迴圈尋找符合條件的 target
	//  */
	// _map := make(map[int]int)
	// index, _len := 0, len(nums)

	// for index < _len {
	// 	another := target - nums[index]
	// 	// 如果 another 已經在 Map 裏面， 代表當前的元素，已經和原本存好的元素配對到了
	// 	if _, ok := _map[another]; ok {
	// 		return []int{_map[another], index}
	// 	}
	// 	// 將 元素、index 存起來
	// 	_map[nums[index]] = index
	// 	index++
	// }

	/** Method 2. (公式)
	 * 先將陣列裡頭元素進行排序，再使用 Two Pointer 方式查找
	 */
	sort.Ints(nums)
	start, end := 0, len(nums)-1
	for start < end {
		sum := nums[start] + nums[end]
		if sum < target { // 代表太小，start 要往右偏移
			start++
		} else if sum > target { // 代表太大，end 要往左偏移
			end--
		} else if sum == target { // 找到囉
			return []int{nums[start], nums[end]}
		}
	}
	return nil
}

// @lc code=end
