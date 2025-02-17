# Intuition
這題主要是要將 nums[i] 處理成 `其餘元素的乘積` (不包含自己 nums[i])，其中時間複雜度只能 `O(n)`，且 `不能使用除法`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 因為不能只用除法，且限制只能 O(n)，只能宣告陣列先 `做左方乘積的運算`，再 `做右方乘積的運算`
- E.g. [1,2,3,4]
    1. `左方乘積`，可以發現 arrLeft[0] 缺少 `[1, 3] 的乘積`，arrLeft[1] 缺少 `[2, 3] 的乘積`，arrLeft[2] 缺少 `[3, 3] 的乘積`，arrLeft[3] 則不缺
    ```
    arrLeft = [1, 1, 2, 6]
    arrLeft[0]: 左邊沒有元素，故保留預設值 1
    arrLeft[1]: 左邊只有一個 nums[0]，乘積為 1
    arrLeft[2]: nums[0] * nums[1]，乘積為 2，可簡化成 arrLeft[1] * nums[1]
    arrLeft[3]: nums[0] * ... * nums[2]，乘積為 6，可簡化成 arrLeft[2] * nums[2]
    ```
    2. `右方乘積`，可以發現 arrRight[3] 缺少 `[0, 2] 的乘積`，arrRight[2] 缺少 `[0, 1] 的乘積`，arrRight[1] 缺少 `[0, 0] 的乘積`，arrRight[0] 則不缺
    ```
    arrRight = [24, 12, 4, 1]
    arrRight[3]: 右邊沒有元素，故保留預設值 1
    arrRight[2]: 左邊只有一個 nums[3]，乘積為 4
    arrRight[1]: nums[3] * nums[2]，乘積為 12，可簡化成 arrRight[2] * nums[2]
    arrRight[0]: nums[3] * ... * nums[1] 乘積為24，可簡化成 arrRight[1] * nums[1]
    ```
    3. 可以發現剛剛處理的兩個陣列彼此是 `互補的`，直接將 arrLeft 與 arrRight `彼此相同 index 的元素進行乘積`，就可以得出答案，意思為 `左邊乘積 arrLeft[x] * 右邊乘積 arrRight[x] = 除了自己其他元素的乘積 ans[x]`
        - 除了自己其他元素的乘積: ans[x]
        - 左邊乘積: arrLeft[x] = nums[0] * ... * nums[x-1]
        - 右邊乘積: arrRight[x] = ans[x+1] * ... * ans[n-1]
    ```
    ans[0] = arrLeft[0] * arrRight[0]
    ans[1] = arrLeft[1] * arrRight[1]
    ans[2] = arrLeft[2] * arrRight[2]
    ans[3] = arrLeft[3] * arrRight[3]
    ```

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

import "sort"

func productExceptSelf(nums []int) []int {
    n := len(nums)
    arr := make([]int, n)
    arr[0] = 1
    for i := 1; i < n; i++ {
        arr[i] = nums[i-1] * arr[i-1]
    }

    right := 1
    for i := n-2; i >= 0; i-- {
        right = right * nums[i+1]
        arr[i] *= right
    }
    return arr
}
```

```java
class Solution {
    public int[] productExceptSelf(int[] nums) {
        int n = nums.length;
        int[] arr = new int[n];
        arr[0] = 1;
        for (int i = 1; i < n; i++) {
            arr[i] = arr[i - 1] * nums[i - 1];
        }

        int right = 1;
        for (int i = n - 2; i >= 0; i--) {
            right *= nums[i + 1];
            arr[i] *= right;
        }
        return arr;
    }
}
```