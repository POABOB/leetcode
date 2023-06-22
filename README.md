# leetcode

## 目錄

| No.    |  Title  |  Solution  |  Acceptance |  Difficulty |  Frequency |
|:--------:|:--------------------------------------------------------------|:--------:|:--------:|:--------:|:--------:|
|0001|Two Sum|[Go](https://github.com/POABOB/leetcode/tree/main/0001.%20Two%20Sum)|50.1%|Easy||
|0167|Two Sum II - Input Array Is Sorted|[Go](https://github.com/POABOB/leetcode/tree/main/0167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted)|60.0%|Medium||
|15|3Sum|[Go](https://github.com/POABOB/leetcode/tree/main/0015.%203Sum)|32.8%|Medium||

## 解法公式

### nSum 解法

- 設定終止條件，當 `n < 2`，代表沒有元素可以遍歷，返回 `res`。
- 如果 `n = 2`，我們可以使用 `Tow Pointer` 的方式來解決。
- 如果 `n > 2`，在每次迴圈中找出 `target - 當前元素`，之後使用遞迴的方式執行 `nSum(nums, n - 1, i + 1, target)`，直到 `n = 2` 的時候，找回符合的元素。

```go
import "sort"

func Sum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 4, 0, target)
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