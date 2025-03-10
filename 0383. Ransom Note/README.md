# Intuition

這題會提供 `ransomNote` 和 `magazine` 兩字串，而 `magazine` 是由 `ransomNote` 內的字元所組成，基本上可以在 Hashmap 上紀錄 `ransomNote` 字元出現次數，逐一比較。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 定義 HashMap 來紀錄 `ransomNote` 字元出現次數，`key` 為 `字元`，`value` 為 `次數`，並定義 `全域計數 count` 用來直接比較用的
- 遍歷 `magazine`
    - 每次遇到 `ransomNote` 字元存在，且字元的計數 > 0，那就把該 value - 1，`全域計數 count` 也扣 1
    - 如果 `全域計數 count` 為 0，代表 `magazine` 是由 `ransomNote` 的字元組成， return true

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

func canConstruct(ransomNote string, magazine string) bool {
	hashmap := make(map[byte]int, 26)
	count := 0
	for k, _ := range ransomNote {
		hashmap[ransomNote[k]]++
		count++
	}

	for k, _ := range magazine {
		if num, ok := hashmap[magazine[k]]; ok && num > 0 {
			hashmap[magazine[k]]--
			count--
		}
		if count == 0 {
			return true
		}
	}
	return false
}
```

```java
class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        Map<Character, Integer> hashmap = new HashMap<>(26);
        int count = 0;
        for (int i = 0; i < ransomNote.length(); i++) {
            hashmap.put(ransomNote.charAt(i), hashmap.getOrDefault(ransomNote.charAt(i), 0) + 1);
            count++;
        }

        for (int i = 0; i < magazine.length(); i++) {
            if (hashmap.containsKey(magazine.charAt(i)) && hashmap.get(magazine.charAt(i)) > 0) {
                hashmap.put(magazine.charAt(i), hashmap.get(magazine.charAt(i)) - 1);
                count--;
            }
            if (count == 0) {
                return true;
            }
        }
        return count == 0;
    }
}
```