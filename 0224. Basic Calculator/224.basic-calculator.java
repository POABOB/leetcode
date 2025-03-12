//Given a string s representing a valid expression, implement a basic
//calculator to evaluate it, and return the result of the evaluation.
//
// Note: You are not allowed to use any built-in function which evaluates
//strings as mathematical expressions, such as eval().
//
//
// Example 1:
//
//
//Input: s = "1 + 1"
//Output: 2
//
//
// Example 2:
//
//
//Input: s = " 2-1 + 2 "
//Output: 3
//
//
// Example 3:
//
//
//Input: s = "(1+(4+5+2)-3)+(6+8)"
//Output: 23
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 3 * 10⁵
// s consists of digits, '+', '-', '(', ')', and ' '.
// s represents a valid expression.
// '+' is not used as a unary operation (i.e., "+1" and "+(2 + 3)" is invalid).
//
// '-' could be used as a unary operation (i.e., "-1" and "-(2 + 3)" is valid).
//
// There will be no two consecutive operators in the input.
// Every number and running calculation will fit in a signed 32-bit integer.
//
//
// Related Topics Math String Stack Recursion 👍 6548 👎 531


import java.util.Stack;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    private String proccessString;

    public int calculate(String s) {
        proccessString = s;
        return fn();
    }

    private int fn() {
        Stack<Integer> stk = new Stack<>();
        int num = 0;
        char op = '+';
        while (!proccessString.isEmpty()) {
            char element = proccessString.charAt(0);
            proccessString = proccessString.substring(1);

            if (element == '(') {
                num = fn();
            }
            if (Character.isDigit(element)) {
                num = num * 10 + Character.getNumericValue(element);
            }

            if ((!Character.isDigit(element) && element != ' ') || proccessString.isEmpty()) {
                int last;
                switch (op) {
                    case '+':
                        stk.push(num);
                        break;
                    case '-':
                        stk.push(-num);
                        break;
                    case '*':
                        last = stk.pop();
                        stk.push(last * num);
                        break;
                    case '/':
                        last = stk.pop();
                        stk.push(last / num);
                        break;
                }
                num = 0;
                op = element;
            }
            if (element == ')') {
                break;
            }
        }

        int sum = 0;
        while (!stk.isEmpty()) {
            sum += stk.pop();
        }
        return sum;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
