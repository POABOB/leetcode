# Intuition

這題的目標是將字串陣列 strs 中的 `異位詞 (anagrams) 分組`，例如：

```
Input:  ["eat", "tea", "tan", "ate", "nat", "bat"]
Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
```
`異位詞 (anagram)`：`字母組成完全相同`但`順序不同`的字串，例如 "eat" 和 "tea"。

- 解題條件
    - 相同字母異序的字串應該被分入同一組。
    - 時間複雜度需要盡量優化。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 為了有效率判斷字串是否為異位詞，我採用 `字母頻率編碼` 來進行哈希映射：
    - 建立 hashmap，`key` 為 `字母頻率的編碼`，`value` 為 `相同頻率對應的字串陣列 index`。
    - 遍歷 strs，將每個字串轉換成 `固定格式的 key`：
        - 使用 []byte 陣列 `統計 a~z 出現的次數`
        - 使用 stringBuilder 將 []byte 轉換成 string 作為 key 來存入 hashmap。
<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(nm)，n 為 `strs 陣列長度`，m 為 `平均字串長度`

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(n)，n 為 `hashmap 儲存 key 的長度`，stringBuilder 不會額外產生空間所以忽略不計

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

func groupAnagrams(strs []string) [][]string {
	answer := make([][]string, 0)
	hashmap := make(map[string]int)
	for _, str := range strs {
		vEncode := encode(str)
		if index, ok := hashmap[vEncode]; ok {
			answer[index] = append(answer[index], str)
		} else {
			hashmap[vEncode] = len(answer)
			answer = append(answer, []string{str})
		}
	}
	return answer
}

func encode(str string) string {
	count := make([]byte, 26)
	for _, v := range str {
		count[v-'a']++
	}
	var builder strings.Builder
	for _, v := range count {
		builder.WriteByte(v + '0') // 轉成字元
	}
	return builder.String()
}

```

```java
import java.util.*;

class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> hashmap = new HashMap<>();

        for (String str : strs) {
            String key = encode(str);
            hashmap.computeIfAbsent(key, k -> new ArrayList<>()).add(str);
        }
        return new ArrayList<>(hashmap.values());
    }

    private String encode(String str) {
        int[] count = new int[26];
        for (char c : str.toCharArray()) {
            count[c - 'a']++;
        }
        return Arrays.toString(count); // 將字母頻率轉成唯一 key
    }
}
```