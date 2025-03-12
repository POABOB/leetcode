package leetcode

import (
	"fmt"
	"unicode"
)

/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */

// @lc code=start
func calculate(s string) int {
	sum, _ := fn(s)
	return sum
}

func fn(s string) (int, string) {
	stk := make([]int, 0)
	var sign byte = '+'
	num := 0
	for len(s) > 0 {
		element := s[0]
		s = s[1:]
		if element == '(' {
			num, s = fn(s)
		}
		if unicode.IsDigit(rune(element)) {
			num = num*10 + int(element-'0')
		}

		if (!unicode.IsDigit(rune(element)) && element != ' ') || len(s) == 0 {
			switch sign {
			case '+':
				stk = append(stk, num)
			case '-':
				stk = append(stk, -num)
			case '*':
				stk[len(stk)-1] *= num
			case '/':
				stk[len(stk)-1] /= num
			}
			sign = element
			num = 0
		}
		if element == ')' {
			break
		}
	}
	sum := 0
	for _, v := range stk {
		sum += v
	}
	return sum, s
}

// @lc code=end
