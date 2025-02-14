# Intuition

在股票市場 prices 中，要如何 `交易k次` 才能夠獲利最大化。這題與 [123. Best Time to Buy and Sell StockIII](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-III/description/) 非常相似。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 找出 DP 的狀態轉移方程式，求出買賣最大值
    - 狀態轉移方程式，這邊的交易次數為 K
      ```go
      // 第 i 天，交易累計 k 次，手上不持有股票
      // 1. 第 i - 1 天，交易累計 k 次，一樣沒持有，持續觀望
      // 2. 第 i - 1 天，交易累計 k 次，持有現價賣出
      dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
      // 第 i 天，交易累計 k 次，手上持有股票
      // 1. 第 i - 1 天，交易累計 k 次，持有但沒有賣出
      // 2. 第 i - 1 天，交易累計 k - 1 次，沒持有，現價買入 (k-1 < 0 代表首次交易，不用抓前值)
      if k-1 < 0 {
          dp[i][k][1] = max(dp[i-1][k][1], -prices[i])
      } else {
          dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
      }
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
    - O(nk)

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(nk)

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

func maxProfit(K int, prices []int) int {
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

	n := len(prices)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, K)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for k := 0; k < K; k++ {
			if i == 0 {
				// 第一天不持有股票的情況，肯定沒有獲利
				dp[i][k][0] = 0
				// 第一天交易持有股票，代表為成本
				dp[i][k][1] = -prices[i]
				continue
			}
			// 第 i 天，交易累計 k 次，手上不持有股票
			// 1. 第 i - 1 天，交易累計 k 次，一樣沒持有，持續觀望
			// 2. 第 i - 1 天，交易累計 k 次，持有現價賣出
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// 第 i 天，交易累計 k 次，手上持有股票
			// 1. 第 i - 1 天，交易累計 k 次，持有但沒有賣出
			// 2. 第 i - 1 天，交易累計 k - 1 次，沒持有，現價買入 (k-1 < 0 代表首次交易，不用抓前值)
			if k-1 < 0 {
				dp[i][k][1] = max(dp[i-1][k][1], -prices[i])
			} else {
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			}
		}
	}
	return dp[n-1][K-1][0]
}
```

```java
class Solution {
    public int maxProfit(int K, int[] prices) {
        int n = prices.length;
        int[][][] dp = new int[n][K][2];
        for (int i = 0; i < n; i++) {
            for (int k = 0; k < K; k++) {
                if (i == 0) {
                    dp[i][k][0] = 0;
                    dp[i][k][1] = -prices[i];
                    continue;
                }
                dp[i][k][0] = Math.max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i]);
                if (k - 1 < 0) {
                    dp[i][k][1] = Math.max(dp[i - 1][k][1], -prices[i]);
                } else {
                    dp[i][k][1] = Math.max(dp[i - 1][k][1], dp[i - 1][k - 1][0] - prices[i]);
                }
            }
        }
        return dp[n - 1][K - 1][0];
    }
}
```