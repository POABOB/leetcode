# Intuition
這題給定兩個字串 s 與 t，要判斷 `s 是否為 t 的子字串`。判斷方式是只要 `s` 的每個字元都在 `t` 中出現 `即可`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 使用 Two Pointer `tIndex` 與 `sIndex` 指針，使用 for 迴圈逐一比對
  - 若是 `t[tIndex] == s[sIndex]` 時，則 `sIndex` 向前移動一個位置
  - 若 `sIndex 等於 s 的長度` 時，則 return true
  - 無論如何，`tIndex` 都要向前移動一個位置
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

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sIndex, tIndex := 0, 0
	for len(t) > tIndex {
		if t[tIndex] == s[sIndex] {
			sIndex++
		}
		tIndex++

		if sIndex == len(s) {
			return true
		}
	}
	return false
}
```

```java
class Solution {
	public boolean isSubsequence(String s, String t) {
		if (s.isEmpty()) {
			return true;
		}
		int sIndex = 0;
		int tIndex = 0;
		while (t.length() > tIndex) {
			if (t.charAt(tIndex) == s.charAt(sIndex)) {
				sIndex++;
			}
			tIndex++;
			if (sIndex == s.length()) {
				return true;
			}
		}
		return false;
	}
}
```