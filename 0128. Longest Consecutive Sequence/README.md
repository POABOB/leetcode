# Intuition

這題就是給定一個 `亂序整數陣列`，希望你可以找出`最長的連續整數長度`，並且要求時間複雜度只能 `O(n)`。

E.g. nums = [100, 1, 2, 4, 3, 200]

可以找到三組分別是 [1, 2, 3, 4], [100], [200]，`最長的長度就是4`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 這題最簡單的方法就是 `使用排序`，然後如果 `發現不連續的元素就重計`。但是排序最低也要花費 `O(nlogn)`
  ，所以只能改用 `空間換時間`。
- 我會先使用 HashMap 紀錄，`key` 放 `陣列元素`，`value` 則是 `紀錄他是否被遍歷`，第一次遍歷我設定成 `false`
  ，第二次則是 `true`。
- 一開始先把所有元素映射成 `false`，然後使用 foreach 隨機遍歷 (`golang map 取元素是隨機的`)。
    - 每次取出來，我會先 `確認 value 是否為 true`，是的話直接 `continue` 代表他被算過了
    - 再來我會以元素鄰近 index 的遞增與遞減，遞增會是 `index, index+1, ... index+n`，當 index+n 不存在 map
      就代表中斷了，遞減則是 `index-1, index-2, ... index-n`，當 index-n 不存在 map 就代表中斷了。
        - 每次遞增/遞減的元素必須把 `value 設定為 true`
    - 最後不斷比較找出最大值

E.g. nums = [100, 1, 2, 4, 3, 200]

假設遍歷第一個元素是 3，那我就會從 3, 4, 5 遞增查找，發現 5 不存在 break，然後從 2, 1, 0 遞減查找，發現 0 不存在 break，得出
count = 4。

如果遍歷第二個元素是 1，因為剛剛遍歷被標記過了，所以會 continue

...

直到 100, 200 被計算完。

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

func longestConsecutive(nums []int) int {
	hashmap := make(map[int]bool)
	for _, v := range nums {
		hashmap[v] = false
	}

	maxCount := 0
	for k, _ := range hashmap {
		if v, _ := hashmap[k]; v {
			continue
		}

		count := 0
		for index := k; ; index++ {
			if _, ok := hashmap[index]; !ok {
				break
			}
			hashmap[index] = true
			count++
		}
		for index := k - 1; ; index-- {
			if _, ok := hashmap[index]; !ok {
				break
			}
			hashmap[index] = true
			count++
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}
```

```java
class Solution {
    public int longestConsecutive(int[] nums) {
        HashMap<Integer, Boolean> map = new HashMap<>();
        for (int num : nums) {
            map.put(num, false);
        }

        int maxCount = 0;
        for (int num : map.keySet()) {
            if (map.get(num)) {
                continue;
            }

            int count = 0;
            for (int i = num; ; i++) {
                if (!map.containsKey(i)) {
                    break;
                }
                map.put(i, true);
                count++;
            }
            for (int i = num - 1; ; i--) {
                if (!map.containsKey(i)) {
                    break;
                }
                map.put(i, true);
                count++;
            }
            maxCount = Math.max(maxCount, count);
        }
        return maxCount;
    }
}
```