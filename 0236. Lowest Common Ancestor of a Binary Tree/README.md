# Intuition

在二元樹中尋找 `兩個節點的最近共同祖先 (Lowest Common Ancestor, LCA)`，可利用遞迴的方式來高效解決這個問題。

# Approach

- 使用遞迴整棵樹，根據以下情況來確定最近共同祖先：
    - 如果 root 為 nil，返回 nil。
    - 如果 `root 等於 p` 或 `q`，則 root 即為 `最近共同祖先`。
    - 遞迴尋找 root.Left 與 root.Right 來 `查找 p 和 q 所在的位置`。
        - 如果 p 和 q 分別位於 `root 的左子樹和右子樹`，則 `root` 即為最近共同祖先。
        - 如果 p 和 q 都在 `同一側子樹`，則只要找到 `最早等於 p 或 q 的節點`。

# Complexity
- Time complexity:
    - O(n)，其中 n 為樹的節點數。
- Space complexity:
    - O(h)，其中 h 為樹的高度，最壞情況下（退化為鏈表）空間複雜度為 O(n)，平均情況下為 O(log n)。

# Code

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
```

```java
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null) {
            return null;
        }
        if (root == p || root == q) {
            return root;
        }

        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);

        if (left == null) {
            return right;
        }
        if (right == null) {
            return left;
        }
        return root;
    }
}
```