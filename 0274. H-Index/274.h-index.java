//Given an array of integers citations where citations[i] is the number of
//citations a researcher received for their iᵗʰ paper, return the researcher's h-index.
//
//
// According to the definition of h-index on Wikipedia: The h-index is defined
//as the maximum value of h such that the given researcher has published at least
//h papers that have each been cited at least h times.
//
//
// Example 1:
//
//
//Input: citations = [3,0,6,1,5]
//Output: 3
//Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each
//of them had received 3, 0, 6, 1, 5 citations respectively.
//Since the researcher has 3 papers with at least 3 citations each and the
//remaining two with no more than 3 citations each, their h-index is 3.
//
//
// Example 2:
//
//
//Input: citations = [1,3,1]
//Output: 1
//
//
//
// Constraints:
//
//
// n == citations.length
// 1 <= n <= 5000
// 0 <= citations[i] <= 1000
//
//
// Related Topics Array Sorting Counting Sort 👍 1564 👎 731


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int hIndex(int[] citations) {
        int n = citations.length;
        int hIndex = 0;
        List<Integer> sortedCitations = Arrays.stream(citations)
                .boxed() // 轉換成 Integer
                .sorted() // 升序排序
                .collect(Collectors.toList());
        for (int i = 1; i <= n; i++) {
            if (sortedCitations.get(n - i) >= i) {
                hIndex = i;
            } else {
                break;
            }
        }
        return hIndex;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
