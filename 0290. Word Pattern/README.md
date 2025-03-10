# Intuition

這題是要把 `字元 pattern` 映射到一個 `字串 word`，可以想像成縮寫，E.g. ty -> thank you，每個 `字串 word` 在 `字串 s` 中會被空格區隔。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 先將 `字串 s` 用空格區隔，存成一個 `字串陣列`，陣列元素為 `word`
- 定義一個 HashMap 與 HashSet，HashMap 的 `key` 為 `pattern`，`value` 為 `word`，HashSet 則是記錄 `word` 是否已經被存過。
    - 如果 HashMap `沒有被存過`，但是 HashSet `有值` 代表 `word` 有被其他 `pattern` 所紀錄，而不是當前的 `pattern`。 
    - 如果 HashMap `有被存過`，與舊映射值不同則 return false
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(n)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	// 紀錄 pattern 對應 word
	patternToWord := make(map[byte]string)
	// Set 紀錄被配過的組合
	wordSet := make(map[string]bool)

	for k, _ := range pattern {
		c := pattern[k]
		w := words[k]

		// 如果不是被收錄在 hasmap
		if _, ok := patternToWord[c]; !ok {
			// 又也被使用過
			if _, ok := wordSet[w]; ok {
				return false
			}
			patternToWord[c] = w
		} else {
			// hasmap 的 word 與原本 word 不同
			if patternToWord[c] != w {
				return false
			}
		}
		wordSet[w] = true
	}
	return true
}
```

```java
class Solution {
    public boolean wordPattern(String pattern, String s) {
        String[] words = s.split(" ");
        if (pattern.length() != words.length) {
            return false;
        }

        Map<Character, String> hashmap = new HashMap<>(26);
        Set<String> hashset = new HashSet<>();
        for (int i = 0; i < pattern.length(); i++) {
            String word = words[i];
            char p = pattern.charAt(i);

            if (!hashmap.containsKey(p)) {
                if (hashset.contains(word)) {
                    return false;
                }
                hashmap.put(p, word);
            } else if (!hashmap.get(p).equals(word)) {
                return false;
            }
            hashset.add(word);
        }
        return true;
    }
}
```