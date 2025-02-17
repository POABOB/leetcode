# Intuition
在陣列中，每個元素代表 `當前位置` 可以 `往右跳幾下`，這題要找出最少可以跳幾下可以達到終點。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 建立 `jump 變數`，記錄 `跳的次數`
- 建立 `farthest 變數`，`當下可以跳的最大值`
- 建立 `end 變數`，作為下次跳之前，`紀錄區間值`
- for 迴圈進行遍歷
    - 每次將當前可以跳幾下與 farthest 進行比較，若 `當前` 比較大，則將它 `賦值給 farthest`
    - 每次比較就可以知道，在 `當下` 和 `之前位置` 可以跳的最大值
    - 如果 `end == i`，代表 `i~end 之間`，已經紀錄其中的 `farthest`，也代表 `要往下跳了`，所以跳了之後，又要 `再度紀錄 end`
      作為最遠能跳的區間 i~end

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

func jump(nums []int) int {
    n := len(nums)
    // 紀錄總共跳了幾下
    jump := 0
    // 紀錄當下與之前能夠跳最遠位置
    end := 0
    // 紀錄跳最多遠
    farthest := 0
    for i := 0; i < n-1; i++ {
        farthest = max(farthest, i+nums[i])
        // 如果已經在能夠跳得最遠位置，則選擇往下跳，重新紀錄下次跳最遠位置
        if end == i {
            jump++
            end = farthest
        }
    }
    return jump
}
```

```java
class Solution {
    public int jump(int[] nums) {
        int end = 0;
        int jump = 0;
        int farthest = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            farthest = Math.max(farthest, i + nums[i]);
            if (end == i) {
                jump++;
                end = farthest;
            }
        }
        return jump;
    }
}
```