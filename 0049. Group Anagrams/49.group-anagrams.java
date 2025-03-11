//Given an array of strings strs, group the anagrams together. You can return
//the answer in any order.
//
//
// Example 1:
//
//
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
//
//
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Explanation:
//
//
// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form
//each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged
//to form each other.
//
//
// Example 2:
//
//
// Input: strs = [""]
//
//
// Output: [[""]]
//
// Example 3:
//
//
// Input: strs = ["a"]
//
//
// Output: [["a"]]
//
//
// Constraints:
//
//
// 1 <= strs.length <= 10⁴
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.
//
//
// Related Topics Array Hash Table String Sorting 👍 20192 👎 671

import java.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> hashmap = new HashMap<>();
        for (String str : strs) {
            String key = encode(str);
            hashmap.computeIfAbsent(key, k -> new ArrayList<>()).add(str);
        }
        return new ArrayList<>(hashmap.values());
    }

    private String encode(String str) {
        int[] count = new int[26];
        for (char c : str.toCharArray()) {
            count[c - 'a']++;
        }
        return Arrays.toString(count); // 將字母頻率轉成唯一 key
    }
}
//leetcode submit region end(Prohibit modification and deletion)
