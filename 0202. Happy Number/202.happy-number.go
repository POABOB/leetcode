package leetcode

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	slow, fast := n, n
	slow = getSum(slow)
	fast = getSum(getSum(fast))

	for slow != fast {
		slow = getSum(slow)
		fast = getSum(getSum(fast))
	}
	return slow == 1
}

func getSum(n int) int {
	var sum int
	for n > 0 {
		remainder := n % 10
		sum += remainder * remainder
		n /= 10
	}
	return sum
}

// @lc code=end
