# Intuition
這題非常基本，就是給你兩個字串 `haystack` 與 `needle`，要找到 `needle` 在 `haystack` 中的第一個出現位置。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 宣告變數 `m` 代表 `haystack` 的長度，`n` 代表 `needle` 的長度
- for 迴圈要設 `m - n 為終止條件`，避免 index 超出陣列
  - 使用 `substring` 來逐一比較
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

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	for i := 0; i <= m-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
```
```java
class Solution {
    public int strStr(String haystack, String needle) {
        int m = haystack.length();
        int n = needle.length();
        for (int i = 0; i <= m - n; i++) {
            if (needle.equals(haystack.substring(i, i + n))) {
                return i;
            }
        }
        return -1;
    }
}
```