//Given an unsorted array of integers nums, return the length of the longest
//consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
//
// Example 1:
//
//
//Input: nums = [100,4,200,1,3,2]
//Output: 4
//Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
//Therefore its length is 4.
//
//
// Example 2:
//
//
//Input: nums = [0,3,7,2,5,8,4,6,0,1]
//Output: 9
//
//
// Example 3:
//
//
//Input: nums = [1,0,1,2]
//Output: 3
//
//
//
// Constraints:
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics Array Hash Table Union Find 👍 21038 👎 1121


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int longestConsecutive(int[] nums) {
        HashMap<Integer, Boolean> map = new HashMap<>();
        for (int num : nums) {
            map.put(num, false);
        }

        int maxCount = 0;
        for (int num : map.keySet()) {
            if (map.get(num)) {
                continue;
            }

            int count = 0;
            for (int i = num; ; i++) {
                if (!map.containsKey(i)) {
                    break;
                }
                map.put(i, true);
                count++;
            }
            for (int i = num - 1; ; i--) {
                if (!map.containsKey(i)) {
                    break;
                }
                map.put(i, true);
                count++;
            }
            maxCount = Math.max(maxCount, count);
        }
        return maxCount;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
