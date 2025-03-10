//Given two strings ransomNote and magazine, return true if ransomNote can be
//constructed by using the letters from magazine and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
//
//
// Example 1:
// Input: ransomNote = "a", magazine = "b"
//Output: false
//
// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
//Output: false
//
// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
//Output: true
//
//
// Constraints:
//
//
// 1 <= ransomNote.length, magazine.length <= 10⁵
// ransomNote and magazine consist of lowercase English letters.
//
//
// Related Topics Hash Table String Counting 👍 5268 👎 522


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        Map<Character, Integer> hashmap = new HashMap<>(26);
        int count = 0;
        for (int i = 0; i < ransomNote.length(); i++) {
            hashmap.put(ransomNote.charAt(i), hashmap.getOrDefault(ransomNote.charAt(i), 0) + 1);
            count++;
        }

        for (int i = 0; i < magazine.length(); i++) {
            if (hashmap.containsKey(magazine.charAt(i)) && hashmap.get(magazine.charAt(i)) > 0) {
                hashmap.put(magazine.charAt(i), hashmap.get(magazine.charAt(i)) - 1);
                count--;
            }
            if (count == 0) {
                return true;
            }
        }
        return count == 0;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
