//Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will
//be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
//
// Note that s may contain leading or trailing spaces or multiple spaces
//between two words. The returned string should only have a single space separating the
//words. Do not include any extra spaces.
//
//
// Example 1:
//
//
//Input: s = "the sky is blue"
//Output: "blue is sky the"
//
//
// Example 2:
//
//
//Input: s = "  hello world  "
//Output: "world hello"
//Explanation: Your reversed string should not contain leading or trailing
//spaces.
//
//
// Example 3:
//
//
//Input: s = "a good   example"
//Output: "example good a"
//Explanation: You need to reduce multiple spaces between two words to a single
//space in the reversed string.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s contains English letters (upper-case and lower-case), digits, and spaces '
//'.
// There is at least one word in s.
//
//
//
// Follow-up: If the string data type is mutable in your language, can you
//solve it in-place with O(1) extra space?
//
// Related Topics Two Pointers String 👍 9139 👎 5318


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public String reverseWords(String s) {
        int p1 = s.length();
        int p2 = s.length();
        StringBuilder sb = new StringBuilder();
        while (p1 >= 0) {
            p1--;
            if (p1 >= 0 && s.charAt(p1) != ' ') {
                continue;
            }

            // 遇到兩個空格，像是"a  a"
            if (p1 + 1 == p2) {
                p2 = p1;
                continue;
            }
            sb.append(" ");
            sb.append(s.substring(p1 + 1, p2));
            p2 = p1;
        }
        return sb.toString().trim();
    }
}
//leetcode submit region end(Prohibit modification and deletion)
