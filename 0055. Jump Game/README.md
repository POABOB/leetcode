# Intuition
在陣列中，每個元素代表 `當前位置` 可以 `往右跳幾下`，所以要確保在跳的時候，是否可以跳到最後面。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 建立 `jump 變數`，記錄每次在跳的時候可以跳得最大值
- for 迴圈進行遍歷
    - 每次將當前可以跳幾下與 jump 進行比較，若 `當前` 比較大，則將它 `賦值給 jump`
    - 每次比較就可以知道，在 `當下` 和 `之前位置` 可以跳的最大值
    - 如果 `i >= jump`，代表走到這個位置無論如何都 `無法前進`，無論之前位置怎麼選，直接 return false
<!-- Describe your approach to solving the problem. -->

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

func canJump(nums []int) bool {
	n := len(nums)
	jump := 0
	for i := 0; i < n-1; i++ {
		jump = max(jump, i+nums[i])
		// 遇到卡住就直接 return，避免錯誤答案 E.g. [0, 2, 3]
		if i >= jump {
			return false
		}
	}
	return jump >= n-1
}
```

```java
class Solution {
    public boolean canJump(int[] nums) {
        int n = nums.length;
        int jump = 0;
        for (int i = 0; i < n - 1; i++) {
            jump = Math.max(jump, i + nums[i]);
            if (i >= jump) {
                return false;
            }
        }
        return jump >= n - 1;
    }
}
```