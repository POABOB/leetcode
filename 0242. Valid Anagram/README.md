# Intuition

這題要檢查兩個字串 `s` 和 `t` 是否為 `字母異位詞 (Anagram)`，也就是 t 是否由 s 重新排列而成。

## 解題條件
- 兩個字串長度需相同，否則必定不是 anagram。
- 兩個字串的字母頻率要完全一致。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 由於字母範圍 僅限於 `小寫英文字母 ('a' - 'z')`，我們可以使用大小為 `26` 的陣列來 `統計字母出現次數`。
- 若 `s.length() != t.length()`，則直接回傳 false。
- 統計字母頻率，用 int[26] 陣列記錄 s 中字母的出現次數。
- 檢查陣列是否全為 0，若 count 陣列所有值都為 0，代表 s 和 t 擁有相同的字母頻率，則回傳 true。

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

func isAnagram(s string, t string) bool {
	sCount := count(s)
	tCount := count(t)
	for i := 0; i < 26; i++ {
		if sCount[i] != tCount[i] {
			return false
		}
	}
	return true
}

func count(s string) []int {
	charCount := make([]int, 26)
	for _, v := range s {
		charCount[v-'a']++
	}
	return charCount
}
```

```java
class Solution {
	public boolean isAnagram(String s, String t) {
		if (s.length() != t.length()) return false;
		int[] sCount = count(s);
		int[] tCount = count(t);
		for (int i = 0; i < 26; i++) {
			if (sCount[i] != tCount[i]) {
				return false;
			}
		}
		return true;
	}

	private int[] count(String s) {
		int[] charCount = new int[26];
		for (char c : s.toCharArray()) {
			charCount[c - 'a']++;
		}
		return charCount;
	}
}
```