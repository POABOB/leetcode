# Intuition

給定一個 `陣列 nums`，透過 `k` 來決定陣列內 `元素要向右 shift 幾次`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- `Method 1.` 最簡單方式，宣告一個一模一樣的陣列，然後將 nums 從 `第 k 位` 到 `第 len(nums) - 1 位 (最後)`
  元素塞入新的陣列，再來從 `第 0 位` 塞到 `第 k - 1 位`
- `Method 2.` 使用特定解法，將 `陣列的元素先進行 reverse`
  ，然後把陣列 `切分成兩個部分 ([0, ..., k - 1], [k, ..., len(nums) - 1])`，`兩陣列獨自進行元素內的 reverse`
    - 假定 nums = [1, 2, 3, 4, 5, 6, 7], k = 3

    1. 陣列元素 reverse: [7, 6, 5, 4, 3, 2, 1]
    2. 切分成兩陣列: [7, 6, 5], [4, 3, 2, 1]
    3. 兩陣列元素獨自 reverse: [7, 6, 5], [4, 3, 2, 1] -> [5, 6, 7], [1, 2, 3, 4]

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - Method 1. O(n)
    - Method 2. O(n)

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - Method 1. O(n)
    - Method 2. O(1)

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

// Method 1. 太基本不寫了
// Method 2. 翻轉, 分割, 獨自翻轉
func rotate(nums []int, k int)  {
    size := len(nums)
    k %= size
    if k != 0 { 
        doRotate(nums, 0, size-1)
        doRotate(nums, 0, k-1)
        doRotate(nums, k, size-1)
    }
}

func doRotate(nums []int, start, end int) {
    for end > start {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
```

```java
class Solution {
    public void rotate(int[] nums, int k) {
        int size = nums.length;
        k = k % size;
        doRatate(nums, 0, size - 1);
        doRatate(nums, 0, k - 1);
        doRatate(nums, k, size - 1);
    }

    public void doRotate(int[] nums, int start, int end) {
        while (end > start) {
            int temp = nums[end];
            nums[end] = nums[start];
            nums[start] = temp;
            end--;
            start++;
        }
    }
}
```
