package leetcode

import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stk := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			stk[len(stk)-2] += stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		case "-":
			stk[len(stk)-2] -= stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		case "/":
			stk[len(stk)-2] /= stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		case "*":
			stk[len(stk)-2] *= stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		default:
			num, _ := strconv.Atoi(token)
			stk = append(stk, num)
		}
	}
	return stk[0]
}

// @lc code=end
