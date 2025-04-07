//Given an integer n, return the number of trailing zeroes in n!.
//
// Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.
//
//
// Example 1:
//
//
//Input: n = 3
//Output: 0
//Explanation: 3! = 6, no trailing zero.
//
//
// Example 2:
//
//
//Input: n = 5
//Output: 1
//Explanation: 5! = 120, one trailing zero.
//
//
// Example 3:
//
//
//Input: n = 0
//Output: 0
//
//
//
// Constraints:
//
//
// 0 <= n <= 10⁴
//
//
//
// Follow up: Could you write a solution that works in logarithmic time
//complexity?
//
// Related Topics Math 👍 3337 👎 1973


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int trailingZeroes(int n) {
        int res = 0;
        int divider = 5;
        while (n >= divider) {
            res += n / divider;
            divider *= 5;
        }
        return res;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
