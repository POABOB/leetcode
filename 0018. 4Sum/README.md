# Intuition
這題思路與 [`0015. 3Sum`](https://github.com/POABOB/leetcode/tree/main/0015.%203Sum) 相同，只不過多一層需要處理，並避免重複。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- `Method 1.` 
	- `第一層迴圈`：`i` 從 `0~(length - 3)` (還有`其他 3 個數`) 遍歷，如果 `nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target` 代表 `i` 有機會找到 `target`，超過就`大於target`
	- `第二層迴圈`：`j` 從 `(i + 1)~(length - 2)` (`不可以與i重複`)，迴圈執行條件與第一層邏輯相同
	- `第三層迴圈`： `Two Sum`，如果 `pointer` 與前一個重複，直接 `++/--` 跳過
- `Method 2.` 
	1. 先將陣列進行排序
	2. 執行 `nSum`
		- 設定終止條件，當 `n < 2`，代表沒有元素可以遍歷，返回 `res`。
		- 如果 `n = 2`，我們可以使用 `Tow Pointer` 的方式來解決。
		- 如果 `n > 2`，在每次迴圈中找出 `target - 當前元素`，之後使用遞迴的方式執行 `nSum(nums, n - 1, i + 1, target)`，直到 `n = 2` 的時候，找回符合的元素。
- 參考資料
    - https://github.com/labuladong/fucking-algorithm/discussions/1021
    - https://github.com/halfrost/LeetCode-Go/blob/master/leetcode/0018.4Sum/18.%204Sum.go
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - Method 1. O(n^3)
    - Method 2. O(n^3)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - Method 1. O(1)
    - Method 2. O(n^3)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	// Method 1.
	result, n := make([][]int, 0), len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result

	// Method 2. 公式 nSum
	// sort.Ints(nums)
	// return nSum(nums, 4, 0, 0)
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
```