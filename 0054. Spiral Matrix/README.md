# Intuition

這題要求按照 `順時針` 螺旋順序遍歷矩陣，即從左上角開始，先 `向右`、`向下`、`向左`、`向上`，依此循環，直到所有元素都被遍歷。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 設定四個邊界變數 `up`、`down`、`left`、`right` 來控制遍歷範圍。
- 使用 `walk 函式` 處理四個方向的遍歷，並在 `每次遍歷後縮小相應的邊界`。
- 持續執行這個過程，直到收集完所有元素。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(mn)，每個元素最多被訪問一次，其中 m 和 n 分別為矩陣的行數與列數。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(1)

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

func spiralOrder(matrix [][]int) []int {
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	total := (right + 1) * (down + 1)
	ans := make([]int, 0, len(matrix)*len(matrix[0]))

	for total > len(ans) {
		ans = walk(matrix, ans, up, down, left, right, 0) // 向右
		up++
		ans = walk(matrix, ans, up, down, left, right, 1) // 向下
		right--
		ans = walk(matrix, ans, up, down, left, right, 2) // 向左
		down--
		ans = walk(matrix, ans, up, down, left, right, 3) // 向上
		left++
	}
	return ans
}

func walk(matrix [][]int, ans []int, up, down, left, right, direction int) []int {
	if len(ans) == len(matrix)*len(matrix[0]) {
		return ans
	}
	if direction == 0 {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
	} else if direction == 1 {
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
	} else if direction == 2 {
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
	} else if direction == 3 {
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
	}
	return ans
}
```

```java
import java.util.ArrayList;
import java.util.List;

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
```