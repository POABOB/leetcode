# Intuition

給定一棵 `二元樹`，檢查它是否是 `對稱的 (鏡像) `。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## 深度優先搜尋 (DFS - 遞迴)

- 如果根節點為 nil，則直接 return true。
- isMirror 遞迴檢查：
    - 若 `兩個節點都為 nil`，則 return true。
    - 若 `其中一個為 nil` 或 `兩者值不同`，則 return false。
    - 若 `兩者值相同`，則繼續比較 `[左節點的左子節點, 右節點的右子節點]` 以及 `[左節點的右子節點, 右節點的左子節點]`。
- 若所有檢查都通過，則 return true。

## 廣度優先搜尋 (BFS - 層序遍歷)

- 使用 `隊列` 來遍歷檢查 `每個節點是否對稱`。
    - 如果根節點為 nil，則直接 return true。
    - 初始化一個隊列，將 [root.Left, root.Right] 放入其中。
    - 進入迴圈，當隊列不為空時：
        - 取出隊列中的 `第一對節點`。
        - 如果 `兩個節點都為 nil`，則繼續下一輪。 (代表節點 `到底了`，而且兩個都到底)
        - 如果 `其中一個為 nil` 或 `兩者的值不同`，則 return false。
        - 將 `[左節點的左子節點, 右節點的右子節點]` 和 `[左節點的右子節點, 右節點的左子節點]` 加入隊列。
    - 如果迴圈結束 return true。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - Method 1. O(n)，每個節點都會被訪問一次。
    - Method 2. O(n)，每個節點都會被訪問一次

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - Method 1. O(h)，該樹的高度，最壞情況是 O(n)，function call stack。
    - Method 2. O(w)，w 為某層的對大節點數，最壞情況是 O(n)。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// Method 1. DFS
	// if root == nil {
	// 	return true
	// }
	// return isMirror(root.Left, root.Right)

	// Method 2. BFS
	queue := make([][2]*TreeNode, 1)
	queue = append(queue, [2]*TreeNode{root.Left, root.Right})
	for len(queue) > 0 {
		leftNode := queue[0][0]
		rightNode := queue[0][1]
		queue = queue[1:]

		if leftNode == nil && rightNode == nil {
			continue
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}

		queue = append(queue, [2]*TreeNode{leftNode.Left, rightNode.Right})
		queue = append(queue, [2]*TreeNode{rightNode.Left, leftNode.Right})
	}
	return true
}

// DFS
// func isMirror(n1 *TreeNode, n2 *TreeNode) bool {
// 	if n1 == nil && n2 == nil {
// 		return true
// 	}
// 	if n1 == nil || n2 == nil {
// 		return false
// 	}
// 	return (n1.Val == n2.Val && isMirror(n1.Left, n2.Right) && isMirror(n1.Right, n2.Left))
// }
```

```java
import java.util.LinkedList;
import java.util.Queue;

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
    public boolean isSymmetric(TreeNode root) {
        // Method 1. DFS
        // if (root == null) return true;
        // return isMirror(root.left, root.right);

        // Method 2. BFS
        if (root == null) return true;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root.left);
        queue.add(root.right);

        while (!queue.isEmpty()) {
            TreeNode n1 = queue.poll();
            TreeNode n2 = queue.poll();

            if (n1 == null && n2 == null) continue;
            if (n1 == null || n2 == null || n1.val != n2.val) return false;

            queue.add(n1.left);
            queue.add(n2.right);
            queue.add(n1.right);
            queue.add(n2.left);
        }
        return true;
    }

    private boolean isMirror(TreeNode n1, TreeNode n2) {
        if (n1 == null && n2 == null) return true;
        if (n1 == null || n2 == null) return false;
        return (n1.val == n2.val) && isMirror(n1.left, n2.right) && isMirror(n1.right, n2.left);

    }
}
```