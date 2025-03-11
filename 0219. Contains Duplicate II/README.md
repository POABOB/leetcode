# Intuition
這題要判斷在給定的數字陣列中，`是否存在兩個相同的元素`，且它們的`索引差不超過 k`。這是一個典型的 `滑動視窗(Sliding Window)` 問題，可以利用 `hashmap` 來快速檢查元素是否在範圍內出現過。

# Approach

- 使用一個 hashmap，`key` 為 `陣列元素`，`value` 為 `元素 index`
- 遍歷陣列中的每個元素，對於每個元素，檢查它是否在哈希表中已經出現過：
    - 如果有，`索引差距小於或等於 k`，則 return true
    - 如果沒有，將 `該數字` 及 `其索引` 儲存
- 若遍歷結束後都沒有找到符合條件的元素，則 return false。

# Complexity
- Time complexity:
    - O(n)，n 為 `陣列的長度`。
- Space complexity:
    - O(n)

# Code

```go
package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
    hashmap := make(map[int]int)
    for k1, v1 := range nums {
        if k2, ok := hashmap[v1]; ok && k1-k2 <= k {
            return true
        }
        hashmap[v1] = k1
    }
    return false
}
```

```java
import java.util.HashMap;
import java.util.Map;

class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Map<Integer, Integer> hashmap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (hashmap.containsKey(nums[i])) {
                if (i - hashmap.get(nums[i]) <= k) {
                    return true;
                }
            }
            hashmap.put(nums[i], i);
        }
        return false;
    }
}
```