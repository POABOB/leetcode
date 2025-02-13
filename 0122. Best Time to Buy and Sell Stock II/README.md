# Intuition

在股票市場 prices 中，要如何 `交易無限次` 才能夠獲利最大化。這題與 [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/) 非常相似，差別差在交易次數。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 找出 DP 的狀態轉移方程式，求出買賣最大值
    - 狀態轉移方程式，這邊的交易次數並沒有特別限制，代表 k 為正無限，於是不用特別紀錄 k 狀態
      ```go
      // 第 i 天，手上不持有股票
      // 1. 第 i - 1 天，一樣沒持有，持續觀望
      // 2. 第 i - 1 天，持有現價賣出
      dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
      // 第 i 天，手上持有股票
      // 1. 第 i - 1 天，持有，持續觀望
      // 2. 第 i - 1 天，沒持有但現價買入
      dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
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
- 參考資料
    - https://labuladong.online/algo/dynamic-programming/stock-problem-summary/

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

	// 由於這題只能可以交易無限次，所以 K 為 無限，所以不用特別去限制交易次數
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
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第 i 天，手上持有股票
		// 1. 第 i - 1 天，持有，持續觀望
		// 2. 第 i - 1 天，沒持有但現價買入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
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
            dp[i][1] = Math.max(dp[i - 1][1], dp[i - 1][0] - prices[i]);
        }
        return dp[n - 1][0];
    }
}
```