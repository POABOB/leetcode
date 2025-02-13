# Intuition

在股票市場 prices 中，要如何 `交易一次` 才能夠獲利最大化。通常要 `求極值`，非常適合使用 `DP` 來解題，也可以使用一般的解法。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- `Method 1.` 找出 DP 的狀態轉移方程式，求出買賣最大值
    - 狀態轉移方程式
      ```go
      // 第 i 天，手上不持有股票
      // 1. 第 i - 1 天，一樣沒持有，持續觀望
      // 2. 第 i - 1 天，持有現價賣出
      dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
      // 第 i 天，手上持有股票
      // 1. 第 i - 1 天，持有，持續觀望
      // 2. 第 i - 1 天，沒持有但現價買入
      dp[i][1] = max(dp[i-1][1], -prices[i])
      ```
    - Base case
      ```go
      // 第一天的 Base case
      dp[0][0] = 0
      dp[0][1] = -prices[i]
      // 由以下簡化
      //   dp[i][0] 
      // = max(dp[-1][0], dp[-1][1] + prices[i])
      // = max(0, -infinity + prices[i]) = 0
      // = max(dp[-1][0], dp[-1][1] + prices[i])
      // = max(0, -infinity + prices[i]) = 0
      //   dp[i][1] 
      // = max(dp[-1][1], dp[-1][0] - prices[i])
      // = max(-infinity, 0 - prices[i]) 
      // = -prices[i]
      // = max(dp[-1][1], dp[-1][0] - prices[i])
      // = max(-infinity, 0 - prices[i]) 
      // = -prices[i]
      ```
- Method 2. 一般解法，使用迴圈紀錄 `每日最小值 minBuy`，並且與後面的價格進行比價，當日若為最大獲利，就記錄為 `maxProfit`
  ，也就是說，越早買，後面有越多天可以比較。
- 參考資料
    - https://labuladong.online/algo/dynamic-programming/stock-problem-summary/

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

// Method 1. DP
func maxProfit(prices []int) int {
    // dp[i][k][0 or 1]
    // K: 允許買賣次數
    // n: 交易天數
    // 0: 代表不持有股票，1: 代表持有股票
    // 總計 n * K * 2 總狀態
    // 狀態轉移方程式
    // for 0 <= i < n:
    //     for 1 <= k <= K:
    //         for s in {0, 1}:
    //             dp[i][k][s] = max(buy, sell, rest)
    
    // 由於這題只能交易一次，所以 K 為 1
    n := len(prices)
    dp := make([][2]int, n)
    for i := 0; i < n; i++ {
        if i == 0 {
            // 第一天不持有股票的情況，肯定沒有獲利
            dp[i][0] = 0
            // 第一天交易持有股票，代表為成本
            dp[i][1] = -prices[i]
            continue
        }

        // 第 i 天，手上不持有股票
        // 1. 第 i - 1 天，一樣沒持有，持續觀望
        // 2. 第 i - 1 天，持有現價賣出
        dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
        // 第 i 天，手上持有股票
        // 1. 第 i - 1 天，持有，持續觀望
        // 2. 第 i - 1 天，沒持有但現價買入
        dp[i][1] = max(dp[i - 1][1], -prices[i])
    }
    return dp[n - 1][0]
}

// Method 2. 一般解法
func maxProfit(prices []int) int {
    minBuy := 99999
    maxProfix := -99999
    for _, price := range prices {
        if minBuy > price {
            minBuy = price
        }

        if price - minBuy > maxProfix {
            maxProfix = price - minBuy
        }
    }
    return maxProfix
}
```

```java
class Solution {
    public int maxProfit(int[] prices) {
        int n = prices.length;
        int[][] dp = new int[n][2];
        for (int i = 0; i < n; i++) {
            if (i == 0) {
                dp[i][0] = 0;
                dp[i][1] = -prices[i];
                continue;
            }
            dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
            dp[i][1] = Math.max(dp[i - 1][1], -prices[i]);
        }
        return dp[n - 1][0];
    }
}
```