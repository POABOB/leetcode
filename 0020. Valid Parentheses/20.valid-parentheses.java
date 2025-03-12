//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']
//', determine if the input string is valid.
//
// An input string is valid if:
//
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
//
//
// Example 1:
//
//
// Input: s = "()"
//
//
// Output: true
//
// Example 2:
//
//
// Input: s = "()[]{}"
//
//
// Output: true
//
// Example 3:
//
//
// Input: s = "(]"
//
//
// Output: false
//
// Example 4:
//
//
// Input: s = "([])"
//
//
// Output: true
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s consists of parentheses only '()[]{}'.
//
//
// Related Topics String Stack 👍 25373 👎 1858


import java.util.List;
import java.util.Stack;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
	public boolean isValid(String s) {
		Stack<Character> stack = new Stack<>();

		for (int i = 0; i < s.length(); i++) {
			char c = s.charAt(i);
			if (c == '(' || c == '[' || c == '{') {
				stack.push(c);
			} else {
				if (stack.isEmpty()) {
					return false;
				}
				char top = stack.peek();
				if (c == ')' && top == '(' || c == ']' && top == '[' || c == '}' && top == '{') {
					stack.pop();
				} else {
					return false;
				}
			}
		}
		return stack.isEmpty();
	}
}
//leetcode submit region end(Prohibit modification and deletion)
