# Intuition

這題要求 `翻轉一棵二元樹`，即將每個節點的 `左右子樹對調`。可以使用 深度優先搜尋 (DFS) 或 廣度優先搜尋 (BFS) 來解這題。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## 深度優先搜尋 (DFS - 遞迴)

- 若 root == nil，直接 `返回 nil`。
- 遞迴翻轉 `左子樹與右子樹`，並將結果交換。
- 回傳翻轉後的 root。

## 廣度優先搜尋 (BFS - 層序遍歷)

- 使用 隊列 (Queue) 來 `層序遍歷 (Level Order Traversal)` 整棵樹。
- 每次從隊列取出一個節點，`交換其左右子樹`。
- 若節點有 `左或右子樹`，則加入隊列繼續處理。
- 當隊列為空時，回傳 root。

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

func invertTree(root *TreeNode) *TreeNode {
	// Method 1. DFS
	// if root == nil {
	// 	return nil
	// }
	// root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	// return root

	// Method 2. BFS
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
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
    public TreeNode invertTree(TreeNode root) {
        // Method 1. DFS
        // if (root == null) {
        //     return null;
        // }
        // TreeNode temp = root.left;
        // root.left = invertTree(root.right);
        // root.right = invertTree(temp);
        // return root;

        // Method 2. BFS
        if (root == null) {
            return null;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();
            TreeNode temp = node.left;
            node.left = node.right;
            node.right = temp;
            if (node.left != null) {
                queue.add(node.left);
            }
            if (node.right != null) {
                queue.add(node.right);
            }
        }
        return root;
    }
}
```