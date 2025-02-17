# Intuition
這題就是比對陣列中 `每個字串`，找出 `相同前綴`，並返回。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 基本上只能使用兩個 for 迴圈每兩個單字逐一比較
	- 如果 `該字串` 長度比 `第一個字串` 小，直接返回，其為最短前綴
	- 如果在比對中，發現兩者的 `字元不一樣`，直接返回
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n^2)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func longestCommonPrefix(strs []string) string {
	m := len(strs)
	n := len(strs[0])

	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			thisStr, prevStr := strs[row], strs[row-1]
			// 如果長度比 col 小，就不用往後比較了
			// 如果 thisStr[col] != prevStr[col]，找到不同的字元，return
			if col >= len(thisStr) || col >= len(prevStr) || thisStr[col] != prevStr[col] {
				return strs[row][:col]
			}
		}
	}
	return strs[0]
}
```
```java
class Solution {
    public String longestCommonPrefix(String[] strs) {
        int m = strs.length;
        int n = strs[0].length();

        for (int col = 0; col < n; col++) {
            for (int row = 1; row < m; row++) {
                String thisString = strs[row];
                String prevString = strs[row - 1];
                if (col >= thisString.length() || col >= prevString.length() || thisString.charAt(col) != prevString.charAt(col)) {
                    return strs[row].substring(0, col);
                }
            }
        }
        return strs[0];
    }
}
```