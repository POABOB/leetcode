//Given an m x n integer matrix matrix, if an element is 0, set its entire row
//and column to 0's.
//
// You must do it in place.
//
//
// Example 1:
//
//
//Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
//Output: [[1,0,1],[0,0,0],[1,0,1]]
//
//
// Example 2:
//
//
//Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
//Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
//
//
// Constraints:
//
//
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -2³¹ <= matrix[i][j] <= 2³¹ - 1
//
//
//
// Follow up:
//
//
// A straightforward solution using O(mn) space is probably a bad idea.
// A simple improvement uses O(m + n) space, but still not the best solution.
// Could you devise a constant space solution?
//
//
// Related Topics Array Hash Table Matrix 👍 15373 👎 781

import java.util.ArrayList;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public void setZeroes(int[][] matrix) {
        int y = matrix.length, x = matrix[0].length;
        List<List<Integer>> zeros = new ArrayList<>();
        for (int i = 0; i < x; i++) {
            for (int j = 0; j < y; j++) {
                if (matrix[j][i] == 0) {
                    zeros.add(List.of(j, i));
                }
            }
        }

        for (List<Integer> zero : zeros) {
            dfs(zero.get(0), zero.get(1), matrix, new int[]{0, 1}); // right
            dfs(zero.get(0), zero.get(1), matrix, new int[]{0, -1}); // left
            dfs(zero.get(0), zero.get(1), matrix, new int[]{1, 0}); // down
            dfs(zero.get(0), zero.get(1), matrix, new int[]{-1, 0}); // up
        }
    }

    private void dfs(int x, int y, int[][] matrix, int[] direction) {
        if (x < 0 || x >= matrix.length || y < 0 || y >= matrix[0].length) {
            return;
        }
        matrix[x][y] = 0;
        dfs(x + direction[0], y + direction[1], matrix, direction);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
