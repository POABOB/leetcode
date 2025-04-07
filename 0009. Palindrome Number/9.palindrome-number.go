package leetcode

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	temp := x
	reverseX := 0
	for temp > 0 {
		reverseX = reverseX*10 + temp%10
		temp /= 10
	}
	return reverseX == x
}

// @lc code=end
