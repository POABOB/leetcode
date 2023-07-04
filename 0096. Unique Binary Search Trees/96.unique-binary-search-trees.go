package leetcode

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {
	// 這題是一個卡塔蘭數(Catalan Number)的範例
	// when n = 0, dp(0) = 1
	// when n = 1, dp(1) = 1
	// when n = 2, dp(2) = dp(0) * dp(1) + dp(1) * dp(0)
	// when n = 3, dp(3) = dp(0) * dp(2) + dp(1) * dp(1) + dp(2) * dp(0)
	// 所以 dp(n) = dp(0) * dp(n-1) + dp(1) * dp(n-2) + … + dp(n-1) * dp(0)

	// Method 1. 使用 for loop + HashMap 來解決
	// 先建立 n+1 大小的一個整數陣列
	// 再使用雙迴圈把dp(1~n)存起來，最後再返回dp(n)就是結果
	// dp := make([]int, n+1)
	// dp[0], dp[1] = 1, 1
	// for i := 2; i <= n; i++ {
	// 	for root := 1; root <= i; root++ {
	// 		dp[i] += dp[root-1] * dp[i-root]
	// 	}
	// }
	// return dp[n]

	// Method 2. 使用 for loop 來解決
	// 其實卡塔蘭數(Catalan Number)有一個通用的公式，那就是排列組合 C(2n, n)/(n + 1)
	// 也就是再 2n 個數字取 n 個數，再除以 n + 1
	ans := 1
	// 2n! / n! * (2n - n)! = 2n! / n! * n!
	// 所以迴圈從 n+1 開始乘到 2n，並且要從 1 (i - n)開始除到 n (2n - n)
	for i := n + 1; i <= 2*n; i++ {
		ans *= i
		ans /= (i - n)
	}
	return ans / (n + 1)

	// Method 3. 使用遞迴來解決，來計算樹的種類，思路與 Method 1. 類似
	// if n <= 1 {
	// 	return 1
	// }

	// ans := 0
	// for i := 1; i <= n; i++ {
	// 	ans += numTrees(i-1) * numTrees(n-i)
	// }

	// return ans
}

// @lc code=end
