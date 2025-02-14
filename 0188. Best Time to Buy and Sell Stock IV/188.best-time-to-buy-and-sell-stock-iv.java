// You are given an integer array prices where prices[i] is the price of a given
// stock on the iᵗʰ day, and an integer k.
//
//  Find the maximum profit you can achieve. You may complete at most k
// transactions: i.e. you may buy at most k times and sell at most k times.
//
//  Note: You may not engage in multiple transactions simultaneously (i.e., you
// must sell the stock before you buy again).
//
//
//  Example 1:
//
//
// Input: k = 2, prices = [2,4,1]
// Output: 2
// Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit =
// 4-2 = 2.
//
//
//  Example 2:
//
//
// Input: k = 2, prices = [3,2,6,5,0,3]
// Output: 7
// Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit =
// 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit
// = 3 -0 = 3.
//
//
//
//  Constraints:
//
//
//  1 <= k <= 100
//  1 <= prices.length <= 1000
//  0 <= prices[i] <= 1000
//
//
//  Related Topics Array Dynamic Programming 👍 7640 👎 215

// leetcode submit region begin(Prohibit modification and deletion)
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
          dp[i][k][1] =
              Math.max(dp[i - 1][k][1], dp[i - 1][k - 1][0] - prices[i]);
        }
      }
    }
    return dp[n - 1][K - 1][0];
  }
}
// leetcode submit region end(Prohibit modification and deletion)
