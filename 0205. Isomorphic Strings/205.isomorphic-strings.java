//Given two strings s and t, determine if they are isomorphic.
//
// Two strings s and t are isomorphic if the characters in s can be replaced to
//get t.
//
// All occurrences of a character must be replaced with another character while
//preserving the order of characters. No two characters may map to the same
//character, but a character may map to itself.
//
//
// Example 1:
//
//
// Input: s = "egg", t = "add"
//
//
// Output: true
//
// Explanation:
//
// The strings s and t can be made identical by:
//
//
// Mapping 'e' to 'a'.
// Mapping 'g' to 'd'.
//
//
// Example 2:
//
//
// Input: s = "foo", t = "bar"
//
//
// Output: false
//
// Explanation:
//
// The strings s and t can not be made identical as 'o' needs to be mapped to
//both 'a' and 'r'.
//
// Example 3:
//
//
// Input: s = "paper", t = "title"
//
//
// Output: true
//
//
// Constraints:
//
//
// 1 <= s.length <= 5 * 10⁴
// t.length == s.length
// s and t consist of any valid ascii character.
//
//
// Related Topics Hash Table String 👍 9557 👎 2174


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean isIsomorphic(String s, String t) {
        Map<Character, Character> sMap = new HashMap<>(26);
        Map<Character, Character> tMap = new HashMap<>(26);

        for (int i = 0; i < s.length(); i++) {
            char sc = s.charAt(i);
            char tc = t.charAt(i);

            if (!sMap.containsKey(sc)) {
                sMap.put(sc, tc);
            } else if (sMap.get(sc) != tc) {
                return false;
            }
            if (!tMap.containsKey(tc)) {
                tMap.put(tc, sc);
            } else if (tMap.get(tc) != sc) {
                return false;
            }
        }
        return true;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
