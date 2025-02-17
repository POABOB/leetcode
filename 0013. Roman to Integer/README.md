# Intuition
這題要把羅馬數字轉換成整數，其中可能要額外注意一點，如果遇到 `左方羅馬數字 < 右方羅馬數字` 的條件，代表左方羅馬數字為 `負數`。

E.g. IV -> -1 + 5 = 4, IX -> -1 + 10 = 9


<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 建立一個 MAP，根據羅馬數字對應相對應的阿拉伯數字
- for 迴圈要設 `len(s) - 1`，避免 index 超出陣列
	- 判斷左右大小，`左大相加`，`右大相減`
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

func romanToInt(s string) int {
	convertMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := convertMap[s[len(s)-1]]
	for i := 0; i < len(s)-1; i++ {
		before := convertMap[s[i]]
		after := convertMap[s[i+1]]
		if after > before {
			total -= before
		} else {
			total += before
		}
	}
	return total
}
```
```java
class Solution {
    public int romanToInt(String s) {
        Map<String, Integer> convertMap =
                Map.of("I", 1, "V", 5, "X", 10, "L", 50, "C", 100, "D", 500, "M", 1000);
        Integer total = convertMap.get(s.substring(s.length() - 1));
        for (int i = 0; i < s.length() - 1; i++) {
            Integer before = convertMap.get(s.substring(i, i + 1));
            Integer after = convertMap.get(s.substring(i + 1, i + 2));
            if (after > before) {
                total -= before;
            } else {
                total += before;
            }
        }
        return total;
    }
}
```