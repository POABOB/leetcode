# Intuition

這是 Conway 的 Game of Life 問題的解法。這是一個零玩家的遊戲，根據 `當前格子的狀態` 和 `周圍格子的狀態` 來決定 `下一回合每個格子的生命狀態`。

- Conway 的 Game of Life 是一個二維網格上的規則如下：
    - 活細胞 (1)：
        - 如果活細胞周圍有 `2 或 3 個` 活細胞，則該細胞在下一回合 `保持活著`。
        - 如果活細胞周圍有 `少於 2 個` 活細胞，則該細胞在下一回合 `死亡 (孤獨）`。
        - 如果活細胞周圍有 `超過 3 個` 活細胞，則該細胞在下一回合 `死亡 (過度擁擠）`。
    - 死細胞 (0)：
        - 如果死細胞周圍正好有 `3 個活細胞`，則該細胞在下一回合 `復活`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 首先，遍歷矩陣 `根據當前格子的狀態`與 `其鄰居` 來標記需要死亡或復活的格子。
    - 使用 `-1` 來標記那些 `原本是活` 的但 `需要死掉` 的格子，
    - 使用 `-2` 來標記那些 `原本是死` 的但 `需要復活` 的格子。
- 完成這些標記後，`再次遍歷矩陣`，將 `-1 更新為 0`，將 `-2 更新為 1`，得到下一回合的結果。
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(mn)，其中 m 是矩陣的行數，n 是列數。
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(1)
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func gameOfLife(board [][]int) {
	y, x := len(board), len(board[0])
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			alive := getLiveNeighbors(j, i, board)
			if board[i][j] == 1 {
				// -1 標記為要死掉
				if alive < 2 || alive > 3 {
					board[i][j] = -1
				}
			} else {
				// -2 標記為要復活
				if alive == 3 {
					board[i][j] = -2
				}
			}
		}
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == -2 {
				board[i][j] = 1
			}
		}
	}
}

func getLiveNeighbors(x, y int, board [][]int) int {
	neighbors := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	alive := 0
	// 遍歷鄰居
	for _, direction := range neighbors {
		nextX, nextY := x+direction[0], y+direction[1]

		// 確保鄰居在邊界內
		if nextY >= 0 && nextY < len(board) && nextX >= 0 && nextX < len(board[0]) {
			if board[nextY][nextX] == 1 || board[nextY][nextX] == -1 {
				alive++
			}
		}
	}
	return alive
}

// func print(board [][]int){
//     for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			fmt.Print(board[i][j])
// 			fmt.Print(" ")
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// 	fmt.Println()
// }
```

```java
class Solution {
	public void gameOfLife(int[][] board) {
		int y = board.length, x = board[0].length;

		for (int i = 0; i < y; i++) {
			for (int j = 0; j < x; j++) {
				int active = getLiveNeighbors(j, i, board);
				if (board[i][j] == 1) {
					if (active < 2 || active > 3) {
						board[i][j] = -1;
					}
				} else {
					if (active == 3) {
						board[i][j] = -2;
					}
				}
			}
		}

		for (int i = 0; i < y; i++) {
			for (int j = 0; j < x; j++) {
				if (board[i][j] == -1) {
					board[i][j] = 0;
				} else if (board[i][j] == -2) {
					board[i][j] = 1;
				}
			}
		}
	}

	private int getLiveNeighbors(int x, int y, int[][] board) {
		int[][] neighbors = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}};
		int alive = 0;
		for (int[] direction : neighbors) {
			int nextX = x + direction[0];
			int nextY = y + direction[1];
			if (nextY >= 0 && nextY < board.length && nextX >= 0 && nextX < board[0].length) {
				if (board[nextY][nextX] == 1 || board[nextY][nextX] == -1) {
					alive++;
				}
			}
		}
		return alive;
	}
}
```