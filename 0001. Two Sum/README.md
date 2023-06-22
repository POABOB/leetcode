# Intuition
Method 1. 透過 HashMap 儲存 `nums 的個別元素` (`map[元素]index`)・再透過迴圈尋找符合條件的 target

(公式)
Method 2. 先將陣列裡頭元素進行`排序`，再使用 `Two Pointer 方式`查找
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 參考資料
    - https://studygolang.com/articles/27605
    - https://labuladong.github.io/algo/di-ling-zh-bfe1b/yi-ge-fang-894da/
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - Method 1. O(n)
    - Method 2. O(logn) ~ O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - Method 1. O(n)
    - Method 2. O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func twoSum(nums []int, target int) []int {
	// Method 1.
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

	// Method 2. (公式)
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
```