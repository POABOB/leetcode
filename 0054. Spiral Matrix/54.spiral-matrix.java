//Given an m x n matrix, return all elements of the matrix in spiral order.
//
//
// Example 1:
//
//
//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [1,2,3,6,9,8,7,4,5]
//
//
// Example 2:
//
//
//Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
// Constraints:
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics Array Matrix Simulation 👍 15776 👎 1412

import java.util.ArrayList;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int left = 0, right = matrix[0].length - 1;
        int up = 0, down = matrix.length - 1;
        int total = (right + 1) * (down + 1);
        List<Integer> ans = new ArrayList<>();
        while (total > ans.size()) {
            walk(matrix, ans, up, down, left, right, 0);
            up++;
            walk(matrix, ans, up, down, left, right, 1);
            right--;
            walk(matrix, ans, up, down, left, right, 2);
            down--;
            walk(matrix, ans, up, down, left, right, 3);
            left++;
        }
        return ans;
    }

    private List<Integer> walk(int[][] matrix, List<Integer> ans, int up, int down, int left, int right, int direction) {
        if (ans.size() == matrix.length * matrix[0].length) {
            return ans;
        }
        if (direction == 0) {
            for (int i = left; i <= right; i++) {
                ans.add(matrix[up][i]);
            }
        } else if (direction == 1) {
            for (int i = up; i <= down; i++) {
                ans.add(matrix[i][right]);
            }
        } else if (direction == 2) {
            for (int i = right; i >= left; i--) {
                ans.add(matrix[down][i]);
            }
        } else if (direction == 3) {
            for (int i = down; i >= up; i--) {
                ans.add(matrix[i][left]);
            }
        }
        return ans;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
