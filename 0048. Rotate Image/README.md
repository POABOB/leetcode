# Intuition

這題要求將一個 n x n 的二維矩陣 `順時針旋轉 90 度`。要達到這個目標，可以將矩陣的 `轉置` 後，再進行`水平翻轉`。具體來說：
1. 轉置：將矩陣的 `行列互換`。
2. 水平翻轉：將 `每一行的元素進行翻轉`。

> Q：如果改成 `逆時針旋轉 90 度` 呢？
>
> A：其實只要將每一列再翻轉就好，原本兩步變成三步

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
- 轉置：對於矩陣中的每個元素 matrix[i][j]，將其與 matrix[j][i] 進行交換，這樣就達到了矩陣的轉置。
	- 要注意，`第一` 與 `第二個迴圈` 終止條件請用 `j < len(matrix)`，如果 `第二個迴圈` 使用 `j < len(matrix[0])`，那就還是會變成原來的矩陣。
- 水平翻轉：對於每一行，將其元素進行翻轉，使矩陣的順序從左到右反向排列。

<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n^2)，因為矩陣的大小是 n x n，需要遍歷每個元素進行轉置和翻轉。
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
	- O(1)，因為只在原矩陣上進行操作，不需要額外的空間。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

func rotate(matrix [][]int) {
	y := len(matrix)
	// 鏡向
	for i := 0; i < y; i++ {
		for j := i + 1; j < y; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 水平翻轉
	for y, _ := range matrix {
		reverse(matrix[y])
	}
}

func reverse(matrixY []int) {
	left, right := 0, len(matrixY)-1
	for right > left {
		matrixY[left], matrixY[right] = matrixY[right], matrixY[left]
		left++
		right--
	}
}
```
```java
class Solution {
	public void rotate(int[][] matrix) {
		int y = matrix.length;
		for (int i = 0; i < y; i++) {
			for (int j = i + 1; j < y; j++) {
				int temp = matrix[i][j];
				matrix[i][j] = matrix[j][i];
				matrix[j][i] = temp;
			}
		}
		for (int[] matrixY : matrix) {
			reverse(matrixY);
		}
	}

	private void reverse(int[] matrixY) {
		int left = 0, right = matrixY.length - 1;
		while (right > left) {
			int temp = matrixY[left];
			matrixY[left] = matrixY[right];
			matrixY[right] = temp;
			left++;
			right--;
		}
	}
}
```