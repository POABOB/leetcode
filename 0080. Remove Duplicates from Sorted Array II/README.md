# Intuition
這題主要是把陣列中，`出現次數大於 2 次` 的元素進行`置換到陣列底部`，並且 `原本的元素要往前移`，其中限制 `空間複雜度為O(1)`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 可以使用 `prev` 和 `now`，去比較前後元素的差異，當 `now != prev` 計數重置成 `1`，相同就 `count++`
- 只要不要遇到 `出現2次以上` 並且 `now == prev`，`pointer`就會一直累加紀錄當前元素，反之，就會跳過不紀錄(`因為相同元素出現超過兩次，不能紀錄`)

# Complexity
- Time complexity
    - O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func removeDuplicates(nums []int) int {
	count, prev, pointer := 0, -9999, 0

	for _, now := range nums {
		if now != prev {
			count = 1
		} else {
			count++
            if count > 2 {
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
```

```java
class Solution {
	public int removeDuplicates(int[] nums) {
		int count = 0;
		int prev = -999999;
		int pointer = 0;
		for (int now : nums) {
			if (now != prev) {
				count = 1;
			} else {
				count++;
				if (count > 2) {
					prev = now;
					continue;
				}
			}

			prev = now;
			nums[pointer] = now;
			pointer++;
		}
		return pointer;
	}
}
```