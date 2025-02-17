# leetcode

## 目錄

| No.  | Title                                  |                                                      Solution                                                      | Acceptance | Difficulty | Frequency |
|:----:|:---------------------------------------|:------------------------------------------------------------------------------------------------------------------:|:----------:|:----------:|:---------:|
| 0001 | Two Sum                                |                        [Go](https://github.com/POABOB/leetcode/tree/main/0001.%20Two%20Sum)                        |   50.1%    |    Easy    |           |
| 0013 | Roman to Integer                       |               [Go, Java](https://github.com/POABOB/leetcode/tree/main/0013.%20Roman%20to%20Integer)                |   64.0%    |    Easy    |           |
| 0015 | 3Sum                                   |                          [Go](https://github.com/POABOB/leetcode/tree/main/0015.%203Sum)                           |   32.8%    |   Medium   |           |
| 0018 | 4Sum                                   |                          [Go](https://github.com/POABOB/leetcode/tree/main/0018.%204Sum)                           |   36.5%    |   Medium   |           |
| 0021 | Merge Two Sorted Lists                 |              [Go](https://github.com/POABOB/leetcode/tree/main/0021.%20Merge%20Two%20Sorted%20Lists)               |   62.8%    |    Easy    |           |
| 0026 | Remove Duplicates from Sorted Array    |    [Go, Java](https://github.com/POABOB/leetcode/tree/main/0026.%20Remove%20Duplicates%20from%20Sorted%20Array)    |   59.2%    |    Easy    |           |
| 0027 | Remove Element                         |                 [Go, Java](https://github.com/POABOB/leetcode/tree/main/0027.%20Remove%20Element)                  |   53.58%   |    Easy    |           |
| 0080 | Remove Duplicates from Sorted Array II | [Go, Java](https://github.com/POABOB/leetcode/tree/main/0080.%20Remove%20Duplicates%20from%20Sorted%20Array%20II)  |   53.6%    |   Medium   |           |
| 0088 | Merge Sorted Array                     |              [Go, Java](https://github.com/POABOB/leetcode/tree/main/0088.%20Merge%20Sorted%20Array)               |   47.23%   |    Easy    |           |
| 0096 | Unique Binary Search Trees             |            [Go](https://github.com/POABOB/leetcode/tree/main/0096.%20Unique%20Binary%20Search%20Trees)             |   59.9%    |   Medium   |           |
| 0121 | Best Time to Buy and Sell Stock        |    [Go, Java](https://github.com/POABOB/leetcode/tree/main/0121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock)    |   54.7%    |    Easy    |           |
| 0122 | Best Time to Buy and Sell Stock II     | [Go, Java](https://github.com/POABOB/leetcode/tree/main/0122.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20II)  |   68.7%    |   Medium   |           |
| 0123 | Best Time to Buy and Sell Stock III    | [Go, Java](https://github.com/POABOB/leetcode/tree/main/0123.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20III) |   50.1%    |    Hard    |           |
| 0167 | Two Sum II - Input Array Is Sorted     |    [Go](https://github.com/POABOB/leetcode/tree/main/0167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted)     |   60.0%    |   Medium   |           |
| 0169 | Majority Element                       |                [Go, Java](https://github.com/POABOB/leetcode/tree/main/0169.%20Majority%20Element)                 |   46.0%    |   Medium   |           |
| 0188 | Best Time to Buy and Sell Stock IV     | [Go, Java](https://github.com/POABOB/leetcode/tree/main/0188.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20IV)  |   45.8%    |    Hard    |           |
| 0189 | Rotate Array                           |                  [Go, Java](https://github.com/POABOB/leetcode/tree/main/0189.%20Rotate%20Array)                   |   39.56%   |   Medium   |           |
| 0274 | H-Index                                |                         [Go](https://github.com/POABOB/leetcode/tree/main/0274.%20H-Index)                         |   38.4%    |   Medium   |           |
| 0380 | Insert Delete GetRandom O(1)           |           [Go](https://github.com/POABOB/leetcode/tree/main/0380.%20Insert%20Delete%20GetRandom%20O(1))            |   52.8%    |   Medium   |           |

## 解法公式

### nSum 解法

- 設定終止條件，當 `n < 2`，代表沒有元素可以遍歷，返回 `res`。
- 如果 `n = 2`，我們可以使用 `Tow Pointer` 的方式來解決。
- 如果 `n > 2`，在每次迴圈中找出 `target - 當前元素`，之後使用遞迴的方式執行 `nSum(nums, n - 1, i + 1, target)`
  ，直到 `n = 2` 的時候，找回符合的元素。

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