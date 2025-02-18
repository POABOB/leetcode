# Intuition
這題會給一個字串，該字串是會 `由空白區隔出不同單字`，需要 `將單字的排序給 reverse`，而 `不能 reverse 字元`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- Method 1. 把字串根據空格切割成多筆至陣列中，然後再陸續 append 到新的陣列，最後重新組裝成字串
- Method 2. 把字串根據空格切割成多筆至陣列中，將該陣列元素 reverse，最後重新組裝成字串
- Method 3. 使用 Tow Pointer 從尾部慢慢組裝，配合 String Builder 節省記憶體空間
  - 如果 `s[p1] 遇到空白`，要停下來 `紀錄 s[p1+1:p2] 的字串`，然後 `p1 = p2`
  - 如果 `s[p1] 遇到空白`，且 p1+1 == p2，代表 `連續 s[p1], s[p2] 都是空白`，`p1 = p2`
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - Method 1. O(n)
    - Method 2. O(n)
    - Method 3. O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - Method 1. O(n)
    - Method 2. O(n)
    - Method 3. O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

import "strings"

func reverseWords(s string) string {

	// Method 1. 把字串根據空格切割成多筆至陣列中，然後再陸續 append 到新的陣列，最後重新組裝成字串
	// words := strings.Fields(s)
	// var ans []string
	// for i := len(words) - 1; i >= 0; i-- {
	// 	if words[i] != "" {
	// 		ans = append(ans, words[i])
	// 	}
	// }
	// return strings.Join(ans, " ")

	// Method 2. 把字串根據空格切割成多筆至陣列中，將該陣列元素 reverse，最後重新組裝成字串
	// words := strings.Fields(s)
	// i, j := 0, len(words) - 1
	// for j > i {
	//     words[i], words[j] = words[j], words[i]
	//     i++
	//     j--
	// }
	// return strings.Join(words, " ")

	// Method 3. 使用 Tow Pointer 從尾部慢慢組裝，配合 String Builder 節省記憶體空間
	var result strings.Builder
	p1, p2 := len(s), len(s)
	for p1 >= 0 {
		p1--
		if p1 >= 0 && s[p1] != ' ' {
			continue
		}
		// 遇到兩個空格，像是"a  a"
		if p1+1 == p2 {
			p2 = p1
			continue
		}
		if len(result.String()) > 0 {
			result.WriteString(" ")
		}

		result.WriteString(s[p1+1 : p2])
		p2 = p1
	}
	return result.String()
}
```
```java
class Solution {
    public String reverseWords(String s) {
        int p1 = s.length();
        int p2 = s.length();
        StringBuilder sb = new StringBuilder();
        while (p1 >= 0) {
            p1--;
            if (p1 >= 0 && s.charAt(p1) != ' ') {
                continue;
            }

            // 遇到兩個空格，像是"a  a"
            if (p1 + 1 == p2) {
                p2 = p1;
                continue;
            }
            sb.append(" ");
            sb.append(s.substring(p1 + 1, p2));
            p2 = p1;
        }
        return sb.toString().trim();
    }
}
```