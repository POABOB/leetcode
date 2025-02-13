package leetcode

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
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
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第 i 天，手上持有股票
		// 1. 第 i - 1 天，持有，持續觀望
		// 2. 第 i - 1 天，沒持有但現價買入
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// @lc code=end
