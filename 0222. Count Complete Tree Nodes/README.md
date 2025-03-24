# Intuition

計算 `完全二元樹的節點數`，根據樹的結構特性，有兩種主要方法來解決這個問題。

# Approach

## Method 1: 遍歷所有節點

- 使用遞迴遍歷樹中的 `每個節點`，並計算總節點數。
    - 如果 root 為 nil，返回 0。
    - 遞迴計算左子樹與右子樹的節點數，加上根節點，即 1 + countNodes(root.Left) + countNodes(root.Right)。

## Method 2: 利用完全二元樹的特性

- 對於 `完全二元樹`，如果 `左子樹的高度` 等於 `右子樹的高度`，則該 `樹的節點數為 2^h - 1`。
    - 如果 root 為 nil，返回 0。
    - 計算 左子樹的高度 `LeftHeight` 與 右子樹的高度 `RightHeight`。
        - 如果 LeftHeight == RightHeight，則 `節點總數為 (1 << LeftHeight) - 1`。
        - 否則，遞迴計算 `1 + countNodes(root.Left) + countNodes(root.Right)`。

# Complexity
- Time complexity:
    - Method 1. O(n)，其中 n 為樹的節點數。
    - Method 2. O(log n * log n)，其中 n 是表達式的長度。每個字符最多被處理一次。
- Space complexity:
    - Method 1. O(n)，其中 n 為樹的節點數，call stack。
    - Method 2. O(log n)，遞迴深度最多為 log n，call stack。

# Code

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	// Method 1. 全部走過就是樹的節點數
	// if root == nil {
	// 	return 0
	// }

	// left := countNodes(root.Left) + 1
	// right := countNodes(root.Right) + 1

	// return left + right - 1 // root 重複

	// Method 2. 完全二元樹節點數為 2^h - 1
	if root == nil {
		return 0
	}
	tempLeft, tempRight := root, root
	LeftHeight, RightHeight := 0, 0
	for tempLeft != nil {
		tempLeft = tempLeft.Left
		LeftHeight++
	}
	for tempRight != nil {
		tempRight = tempRight.Right
		RightHeight++
	}
	if LeftHeight == RightHeight {
		return (1 << LeftHeight) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)

}
```

```java
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    public int countNodes(TreeNode root) {
        if (root == null) {
            return 0;
        }
        TreeNode left = root;
        TreeNode right = root;
        int leftHeight = 0;
        int rightHeight = 0;
        while (left != null) {
            left = left.left;
            leftHeight++;
        }
        while (right != null) {
            right = right.right;
            rightHeight++;
        }
        if (leftHeight == rightHeight) {
            return (1 << leftHeight) - 1;
        }
        return 1 + countNodes(root.left) + countNodes(root.right);
    }
}
```