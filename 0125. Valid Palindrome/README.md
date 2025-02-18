# Intuition
這題說的 `Palindrome 回文`，就是將字串切割一半，`後半段字串 reverse 會等於前半段字串`，其中，`標點符號`、`空格` 不算在字串內，`數字` 則要算。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 使用 Two Pointer `left` 與 `right` 指針，使用 for 迴圈前後逐一比對
	- 處理前先將所有字元 `轉小寫`
	- 要封裝判斷 `是否為字元`、`是否為數字` 的函數，如果不是那就是 `該指針前進/後退`
	- 如果當前 `字元不相等`，直接 return false
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

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for right > left {
		if !isLetterOrDigital(rune(s[left])) {
			left++
			continue
		} else if !isLetterOrDigital(rune(s[right])) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetterOrDigital(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsDigit(char)
}
```

```java
class Solution {
    public boolean isPalindrome(String s) {
        s = s.toLowerCase();
        int left = 0;
        int right = s.length() - 1;
        while (right > left) {
            if (!Character.isLetterOrDigit(s.charAt(left))) {
                left++;
                continue;
            } else if (!Character.isLetterOrDigit(s.charAt(right))) {
                right--;
                continue;
            }
            
            if (s.charAt(left) != s.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
}
```