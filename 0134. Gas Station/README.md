# Intuition
這題提供兩個陣列，陣列的元素是要到達的地點，而每個地點都有 `加油站 gas` 與 `耗油量 cost`，為了能夠省油錢將所有地點走過，需要找一個 `最適合出發的起點`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 所謂最適合出發的起點，也就是 `在最耗油的起點` 出發，這樣途中累積油量，就不會突然需要大量油而無法前進
- 一開始需要先紀錄 `累計油量 sum` 與 `最低需要油量 minSum`，只要 `minSum > sum`，代表前往 `第 i + 1 點時油不夠`，那就將 i + 1 設為 `起始點 start`，繼續跑。
- 其中，sum 賦值給 minSum，主要用來判斷從 i + 1 點出發，有沒有辦法 `避免前往下一站沒油`。
- 最後判斷有沒有 `累計油量 sum` 在整體旅程是否為 `正`，反之，則代表 `無論哪一點` 出發都無法成功完成旅程。

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

func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    sum := 0
    minSum := 0
    start := 0
    for i := 0; i < n; i++ {
        sum += gas[i] - cost[i]
        // 行駛途中發現油不夠前往下一站，要換地點上車
        if minSum > sum {
            // 在最耗油的一站開始，才會是最佳解
            start = i + 1
            minSum = sum
        }
    }

    // 到最後油量都是負的
    if sum < 0 {
        return -1
    }

    if start == n {
        return 0
    }
    return start
}
```

```java
class Solution {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int n = gas.length;
        int sum = 0;
        int minSum = 0;
        int start = 0;
        for (int i = 0; i < n; i++) {
            sum += gas[i] - cost[i];
            if (minSum > sum) {
                start = i + 1;
                minSum = sum;
            }
        }

        if (sum < 0) {
            return -1;
        }
        return start % n;
    }
}
```