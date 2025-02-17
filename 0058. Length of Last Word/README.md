# Intuition
這題要將一段字串中，找出 `最後一個單字`，並 `返回該單字長度`，其中字串會由 `空格` 區隔。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 可以使用 Two Pointer `left`, `right`，由字串 `最後一位往前計算`
	- `left 與 right 都是空格`，那就代表 `還沒有找到字`，兩者都往後
	- 只有 `left 是空格` 代表 `該字已經結束`，直接 `break`
	- `left 與 right 都是字元`，`left-- 繼續找`

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

func lengthOfLastWord(s string) int {
	left, right := len(s)-1, len(s)-1
	for left >= 0 {
		if s[left] == ' ' && s[right] == ' ' {
			left--
			right--
		} else if s[left] == ' ' {
			break
		} else {
			left--
		}
	}
	return right - left
}
```

```java
class Solution {
	public int lengthOfLastWord(String s) {
		int left = s.length() - 1;
		int right = s.length() - 1;
		while (left >= 0) {
			if (s.charAt(left) == ' ' && s.charAt(right) == ' ') {
				left--;
				right--;
			} else if (s.charAt(left) == ' ') {
				break;
			} else {
				left--;
			}
		}
		return right - left;
	}
}
```