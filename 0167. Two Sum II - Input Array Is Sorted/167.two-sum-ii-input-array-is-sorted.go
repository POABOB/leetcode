package leetcode

/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	// /** Method 1.
	//  * 透過 HashMap 儲存 nums 的個別元素 (map[元素]index)
	//  * 再透過迴圈尋找符合條件的 target
	//  */
	// _map := make(map[int]int)
	// index, _len := 0, len(numbers)

	// for index < _len {
	// 	another := target - numbers[index]
	// 	// 如果 another 已經在 Map 裏面， 代表當前的元素，已經和原本存好的元素配對到了
	// 	if _, ok := _map[another]; ok {
	// 		return []int{_map[another] + 1, index + 1}
	// 	}
	// 	// 將 元素、index 存起來
	// 	_map[numbers[index]] = index
	// 	index++
	// }

	/** Method 2. (公式)
	 * 不用排序，再使用 Two Pointer 方式查找
	 */
	start, end := 0, len(numbers)-1
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum < target { // 代表太小，start 要往右偏移
			start++
		} else if sum > target { // 代表太大，end 要往左偏移
			end--
		} else if sum == target { // 找到囉
			return []int{start + 1, end + 1}
		}
	}
	return nil
}

// @lc code=end
