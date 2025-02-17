# Intuition
H-Index 定義：h 篇論文中，每篇論文至少有 h 個引用，並且 n(全部論文數) - h 篇論文少於 h 個引用，所以只要將論文的引用數 `進行排序`，由大到小去 `比較該引用數是否有至少 h 個引用`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 先將 citation 升序排序
- 使用 for 迴圈遍歷排查論文 (引用數: 多->少)
	- 如果 `引用數 >= i`，則將 `hIndex 記錄為 i` ，代表 `n 篇論文中至少 i 個引用`
	- 如果 `引用數 < i`，則 `結束運算`。
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

import "sort"

func hIndex(citations []int) int {
    // 定義
    // h papers each have at least h citations
    // And n - h papers have less than h citations
    // E.g.
    // Origin: 3 0 6 1 5
    // Sorted: 0 1 3 5 6
    // 1 篇 paper 至少有 1 次 ，citations[5 - 1] = 6
    // 2 篇 paper 至少有 2 次 ，citations[5 - 2] = 5
    // 3 篇 paper 至少有 3 次 ，citations[5 - 3] = 3
    // 4 篇 paper 至少有 4 次 ，citations[5 - 4] = 1
    // 5 篇 paper 至少有 5 次 ，citations[5 - 5] = 0
    all := len(citations)
    sort.Ints(citations)
    hIndex := 0
    for i := 1; i <= all; i++  {
        if citations[all - i] >= i {
			hIndex = i
		} else {
			break
		}
    }
    return hIndex
}
```

```java
class Solution {
    public int hIndex(int[] citations) {
        int n = citations.length;
        int hIndex = 0;
        List<Integer> sortedCitations = Arrays.stream(citations)
                .boxed() // 轉換成 Integer
                .sorted() // 升序排序
                .collect(Collectors.toList());
        for (int i = 1; i <= n; i++) {
            if (sortedCitations.get(n - i) >= i) {
                hIndex = i;
            } else {
                break;
            }
        }
        return hIndex;
    }
}
```