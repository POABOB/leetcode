# Intuition
這題主要是要刪除 `陣列 nums` 內的 `指定元素 val`，其中要把刪除後的 `空位移動至陣列尾部`，也就是說刪除元素後，要 `把沒被刪除的元素向前移動`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 可以使用 `Two Pointer` 的方式來解決， slow pointer 為 `替換的指標`， fast pointer 為 `遍歷 nums 的指標`
- 遍歷時，要確定 `len(nums) > fast pointer`
  - 當 `nums[fast] != val` 時，就可以把 `nums[fast]` 替換到 `nums[slow]`，然後 `slow pointer + 1`
  - if 條件外一直都要 `fast pointer + 1`
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

func removeElement(nums []int, val int) int {
    fast, slow := 0, 0
    for len(nums) > fast {
        if nums[fast] != val{
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return slow
}
```

```java
class Solution {
    public int removeElement(int[] nums, int val) {
        int fast = 0;
        int slow = 0;
        while (nums.length > fast) {
            if (nums[fast] != val) {
                nums[slow] = nums[fast];
                slow++;
            }
            fast++;
        }
        return slow;
    }
}
```




