# Intuition
這題要求我們在給定的矩陣中將所有 `包含 0 的行與列設為 0`。這是一個典型的矩陣遍歷問題，可以通過在遍歷矩陣時 `標記 0 的位置`，並進行後續的行列操作來解決。
- 可以想像成爆爆王水球爆炸的衝擊波，這樣似乎比較好理解

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 首先，遍歷矩陣，記錄所有為 0 的位置。
- 然後，對於每個記錄的 0 位置，使用 `深度優先搜尋 (DFS)` 來將其所在行和列的所有元素設為 0。
- 由於 DFS 是一個遞迴過程，因此我們會針對每個 0 的位置進行 `四個方向 (上、下、左、右)` 遍歷，將符合條件的元素設為 0。
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity:
	- O(m * n)，其中 m 是矩陣的行數，n 是矩陣的列數。
- Space complexity:
	- O(m * n)，需要額外的空間來儲存所有 0 位置的資訊（在 zeros 切片中）。
# Code
```go
package leetcode

func setZeroes(matrix [][]int) {
	// 先記錄 0 的位置
	y, x := len(matrix), len(matrix[0])
	zeros := make([][]int, 0)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}

	// 使用 dfs 將有 0 的位置的上下左右全部設定為 0
	for _, zeroPosition := range zeros {
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{0, 1})  // 向右
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{0, -1}) // 向左
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{1, 0})  // 向下
		dfs(zeroPosition[0], zeroPosition[1], matrix, []int{-1, 0}) // 向上
	}
}

func dfs(x, y int, matrix [][]int, direction []int) {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
		return
	}
	matrix[x][y] = 0
	dfs(x+direction[0], y+direction[1], matrix, direction)
}
```

```java
import java.util.ArrayList;
import java.util.List;

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
```