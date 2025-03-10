//Given two strings s and t, return true if t is an anagram of s, and false
//otherwise.
//
//
// Example 1:
//
//
// Input: s = "anagram", t = "nagaram"
//
//
// Output: true
//
// Example 2:
//
//
// Input: s = "rat", t = "car"
//
//
// Output: false
//
//
// Constraints:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s and t consist of lowercase English letters.
//
//
//
// Follow up: What if the inputs contain Unicode characters? How would you
//adapt your solution to such a case?
//
// Related Topics Hash Table String Sorting 👍 12798 👎 423


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;
        int[] sCount = count(s);
        int[] tCount = count(t);
        for (int i = 0; i < 26; i++) {
            if (sCount[i] != tCount[i]) {
                return false;
            }
        }
        return true;
    }

    private int[] count(String s) {
        int[] charCount = new int[26];
        for (char c : s.toCharArray()) {
            charCount[c - 'a']++;
        }
        return charCount;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
