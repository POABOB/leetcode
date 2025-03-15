# Intuition

這題要求計算 `二元樹的最大深度`，也就是算出 `樹的高度`。這題使用 DFS 或 BFS 都可以， DFS 相對會比較簡潔。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## 深度優先搜尋 (DFS - 遞迴)

- 如果節點為 nil，返回 0，表示 `樹的深度為 0`。
- 遞迴計算 `左子樹的最大深度` maxDepth(root.Left)。
- 遞迴計算 `右子樹的最大深度` maxDepth(root.Right)。
- 取 `兩者的最大值`，然後 `加 1` (包含當前節點)。

## 廣度優先搜尋 (BFS - 層序遍歷)

- 使用 `隊列 (Queue)` 來進行 `層序遍歷`，每遍歷一層就將 `level +1`。
- 先將 `根節點加入隊列`，每次取出 `當前層所有的節點`，並將其 `子節點加入隊列`。
- 當 `隊列清空` 時，level 就是 `最大深度`。

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

func maxDepth(root *TreeNode) int {
	// Method 1. DFS
	// if root == nil {
	// 	return 0
	// }
	// return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

	// Method 2. BFS
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelLength:]
		level++
	}
	return level
}

```

```java
import java.util.LinkedList;
import java.util.Queue;

class Solution {
    public int maxDepth(TreeNode root) {
        // Method 1. DFS
        // if (root == null) {
        //     return 0;
        // }
        // return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;

        // Method 2. BFS
        if (root == null) {
            return 0;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        int level = 0;
        while (!queue.isEmpty()) {
            int width = queue.size();
            for (int i = 0; i < width; i++) {
                TreeNode node = queue.poll();
                if (node.left != null) {
                    queue.add(node.left);
                }
                if (node.right != null) {
                    queue.add(node.right);
                }
            }
            level++;
        }
        return level;
    }
}
```