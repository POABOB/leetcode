# Intuition

這題給定一個 9x9 的數獨棋盤，要求檢查該數獨棋盤是否有效。有效的條件是：每一行、每一列以及每一個 3x3 的子方格中，數字必須從 1 到 9 唯一且不重複。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 遍歷 `每一行`、`每一列` 以及 `每個 3x3 方格`：
	- 橫向檢查：對於每一行，檢查是否有重複的數字，跳過 '.'。
	- 縱向檢查：對於每一列，檢查是否有重複的數字，跳過 '.'。
	- 3x3 方格檢查：對於每一個 3x3 的方格，檢查是否有重複的數字，跳過 '.'。
- 使用 HashSet 來檢查是否有 `重複的數字`，對每一個有效位置的數字進行檢查。

<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(1)，因為數獨棋盤的大小是固定的 9x9，所有的檢查操作都是對固定大小的數據進行的，與輸入的大小無關。
	- 如果數獨改成 n * n 的格式，那就會是 O(n)
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(1)，因為 HashSet 中最多只會包含 9 個數字（1 到 9），這是一個固定的數量，空間複雜度是常數級別。
	- 如果數獨改成 n * n 的格式，那就會是 O(n)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func isValidSudoku(board [][]byte) bool {
	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		verticle := validDirection(v, v, board, 1)   // 驗證縱向
		horizontal := validDirection(v, v, board, 0) // 驗證橫向
		if !verticle || !horizontal {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !valid3x3(j*3, i*3, board) {
				return false
			}
		}
	}
	return true
}

func valid3x3(x, y int, board [][]byte) bool {
	hashset := make(map[byte]struct{}, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[y+i][x+j] == '.' {
				continue
			}
			if _, ok := hashset[board[y+i][x+j]]; ok {
				return false
			}
			hashset[board[y+i][x+j]] = struct{}{}
		}
	}
	return true
}

func validDirection(x, y int, board [][]byte, direction int) bool {
	hashset := make(map[byte]struct{}, 9)
	if direction == 0 { // 橫向
		for i := 0; i < 9; i++ {
			if board[y][i] == '.' {
				continue
			}
			if _, ok := hashset[board[y][i]]; ok {
				return false
			}
			hashset[board[y][i]] = struct{}{}
		}
	} else {
		for i := 0; i < 9; i++ {
			if board[i][x] == '.' {
				continue
			}
			if _, ok := hashset[board[i][x]]; ok {
				return false
			}
			hashset[board[i][x]] = struct{}{}
		}
	}
	return true
}
```
```java
import java.util.HashSet;
import java.util.Set;

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
```