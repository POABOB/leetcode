package leetcode

/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
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

	// 由於這題只能可以交易2次，所以 K 為 2
	n := len(prices)
	K := 2
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
// @lc code=end
