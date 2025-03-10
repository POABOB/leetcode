# Intuition

給定 s 與 t 字串，判斷彼此是不是 `同構字串`。同構字串就是兩字串字元有 `獨自映射的規律`。

E.g. s = "egg", t = "add"

規律：
1. e 映射到 a
2. g 映射到 d

所以兩者為同構字串！

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 使用 HashMap，for loop 遍歷把 s 字串每個字元映射到 t 字串， `key` 為 `s[i]`，`value` 為 `t[i]`
- 一旦，發現 `兩者映射有不同結果`，直接返回 false
- 再來因為有可能字元 `單方面` 在映射的過程中，`兩字串字元相同`，爾後 `該字元又遇到不同字元`，導致誤判情況
    - E.g. s = "badc" t = "baba"
    - `b 映射 b`, `d 又映射 b`，此時只有使用一個 map 會讓該情況發生
- 額外再建立一個 HashMap，反向把 t 字串每個字元映射到 s 字串， `key` 為 `t[i]`，`value` 為 `s[i]`，這樣就可以避免上述情況

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

func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte, 26)
	tMap := make(map[byte]byte, 26)

	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; !ok {
			sMap[s[i]] = t[i]
		}
		if sMap[s[i]] != t[i] {
			return false
		}

		if _, ok := tMap[t[i]]; !ok {
			tMap[t[i]] = s[i]
		}
		if tMap[t[i]] != s[i] {
			return false
		}
	}
	return true
}
```

```java
class Solution {
    public boolean isIsomorphic(String s, String t) {
        Map<Character, Character> sMap = new HashMap<>(26);
        Map<Character, Character> tMap = new HashMap<>(26);

        for (int i = 0; i < s.length(); i++) {
            char sc = s.charAt(i);
            char tc = t.charAt(i);

            if (!sMap.containsKey(sc)) {
                sMap.put(sc, tc);
            } else if (sMap.get(sc) != tc) {
                return false;
            }
            if (!tMap.containsKey(tc)) {
                tMap.put(tc, sc);
            } else if (tMap.get(tc) != sc) {
                return false;
            }
        }
        return true;
    }
}
```
