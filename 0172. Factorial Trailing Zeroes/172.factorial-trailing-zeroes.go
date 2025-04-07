package leetcode

/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
func trailingZeroes(n int) int {
	// n < 10^4
	// 5   的倍數: 5, 10, 15, 20, 25, 30, 35    -> 提供1個 5
	// 25  的倍數: 25, 50, 75, ...              -> 提供2個 5
	// 125 的倍數: 125, 250, 375, ...           -> 提供3個 5

	// E.g. n = 10
	// 10 / 5 = 2，會有 2 個 5
	// 總共 2 個
	// E.g. n = 30
	// 30 / 5  = 6，會有 6 個 5
	// 30 / 25 = 1，會有 1 個 5 (25 會提供 2 個，所以要補回來)
	// 總共 7 個
	// E.g. n = 130
	// 130 / 5 = 26，會有 26 個 5
	// 130 / 25 = 5，會有 5 個 5 (25 會提供 2 個，所以要補回來)
	// 130 / 125 = 1，會有 1 個 5 (125 會提供 3 個，所以要補回來)
	// 總共 32 個
	res := 0
	divisor := 5
	for n >= divisor {
		res += n / divisor
		divisor *= 5
	}
	return res
}

// @lc code=end
