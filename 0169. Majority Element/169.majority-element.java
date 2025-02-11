//Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times.
//You may assume that the majority element always exists in the array.
//
//
// Example 1:
// Input: nums = [3,2,3]
//Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
//Output: 2
//
//
// Constraints:
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//Follow-up: Could you solve the problem in linear time and in
//O(1) space?
//
// Related Topics Array Hash Table Divide and Conquer Sorting Counting 👍 20369
//👎 706


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int majorityElement(int[] nums) {
        int candicate = 0;
        int vote = 0;

        // 先讓競選人彼此的票互相抵銷，直到剩下最後一個候選人
        for (int num : nums) {
            if (candicate == num) {
                vote++;
            } else if (vote == 0) {
                candicate = num;
                vote = 1;
            } else {
                vote--;
            }
        }

        // 統計剩下候選人票數是否過半
        vote = 0;
        for (int num : nums) {
            if (candicate == num) {
                vote++;
            }
        }
        if (vote > nums.length / 2) {
            return candicate;
        }
        return -1;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
