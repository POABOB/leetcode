//Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
//validated according to the following rules:
//
//
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9
//without repetition.
//
//
// Note:
//
//
// A Sudoku board (partially filled) could be valid but is not necessarily
//solvable.
// Only the filled cells need to be validated according to the mentioned rules.
//
//
//
//
// Example 1:
//
//
//Input: board =
//[["5","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//Output: true
//
//
// Example 2:
//
//
//Input: board =
//[["8","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//Output: false
//Explanation: Same as Example 1, except with the 5 in the top left corner
//being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is
//invalid.
//
//
//
// Constraints:
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.
//
//
// Related Topics Array Hash Table Matrix 👍 11317 👎 1188

import java.util.HashSet;
import java.util.Set;

//leetcode submit region begin(Prohibit modification and deletion)
public class Solution {
    public boolean isValidSudoku(char[][] board) {
        // 檢查每一列與每一行
        for (int i = 0; i < 9; i++) {
            if (!validDirection(i, i, board, 0)) {  // 檢查橫向
                return false;
            }
            if (!validDirection(i, i, board, 1)) {  // 檢查縱向
                return false;
            }
        }

        // 檢查每個 3x3 方格
        for (int i = 0; i < 9; i += 3) {
            for (int j = 0; j < 9; j += 3) {
                if (!valid3x3(i, j, board)) {  // 檢查 3x3 方格
                    return false;
                }
            }
        }
        return true;
    }

    private boolean valid3x3(int x, int y, char[][] board) {
        Set<Character> hashSet = new HashSet<>();
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                char num = board[y + i][x + j];
                if (num == '.') {
                    continue;
                }
                if (hashSet.contains(num)) {
                    return false;
                }
                hashSet.add(num);
            }
        }
        return true;
    }

    private boolean validDirection(int x, int y, char[][] board, int direction) {
        HashSet<Character> hashSet = new HashSet<>();
        if (direction == 0) {  // 檢查橫向
            for (int i = 0; i < 9; i++) {
                char num = board[y][i];
                if (num == '.') {
                    continue;
                }
                if (hashSet.contains(num)) {
                    return false;
                }
                hashSet.add(num);
            }
        } else {  // 檢查縱向
            for (int i = 0; i < 9; i++) {
                char num = board[i][x];
                if (num == '.') {
                    continue;
                }
                if (hashSet.contains(num)) {
                    return false;
                }
                hashSet.add(num);
            }
        }
        return true;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
