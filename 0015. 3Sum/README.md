# Intuition
這題就是 [`0001. Two Sum`](https://github.com/POABOB/leetcode/tree/main/0015.%203Sum) 的進階版，暴力解已經不能使用了一定會爆開，參考 `排序後二元搜尋` 方法後，其實我們只需要在 `Two Sum` 外層加一個迴圈，這樣就可以計算到 `三數之和`。

> 其實還需要特別注意到的一點就是，如果 `三數` 也有可能遇到重複的組合，需要特別排除。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- `Method 1.` 
	1. 先將陣列進行排序
	2. 使用迴圈定位 `index` 作為基準
	3. 再使用 `start`、`end` 來在第二個迴圈之中比較與 `index` 的元素之和是否為 `target(0)`
		- 如果 `start`，與遍歷的前一次(`start-1`)元素相同，那就直接 `start++` 不處理
		- 如果 `end`，與遍歷的前一次(`end+1`)元素相同，那就直接 `end--` 不處理
	4. 再來比較總和如果 > `target`，代表要 `end--`，反之 `start++`。如果總和 == `target`，`end--`與`start++`都要
- `Method 2.` 
	1. 先將陣列進行排序
	2. 執行 `nSum`
		- 設定終止條件，當 `n < 2`，代表沒有元素可以遍歷，返回 `res`。
		- 如果 `n = 2`，我們可以使用 `Tow Pointer` 的方式來解決。
		- 如果 `n > 2`，在每次迴圈中找出 `target - 當前元素`，之後使用遞迴的方式執行 `nSum(nums, n - 1, i + 1, target)`，直到 `n = 2` 的時候，找回符合的元素。
- 參考資料
    - https://github.com/labuladong/fucking-algorithm/discussions/1021
    - https://github.com/halfrost/LeetCode-Go/blob/master/leetcode/0015.3Sum/15.%203Sum.go
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - Method 1. O(n^2)
    - Method 2. O(n^2)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - Method 1. O(1)
    - Method 2. O(n^2)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

import "sort"

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
```