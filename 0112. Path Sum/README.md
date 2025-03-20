# Intuition

給定二元樹的 `根節點` 和 `整數 targetSum`，如果從 `根節點` 到 `任何葉節點的路徑存在`，且 `路徑上的節點值總和` 等於 `targetSum`，則返回 `true`。否則，返回 false。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## Method 1. DFS（深度優先搜尋）

- 通過從 `targetSum 減去當前節點的值` 來更新 `剩餘的總和`
- 如果 `當前節點` 是 `葉節點` 且 `剩餘總和為 0`，則 return true
- 否則，`繼續遍歷` 左子樹和右子樹。
- 如果最後 `節點為 nil`，return false

## Method 2. BFS（廣度優先搜尋）

- 使用 `queue` 來遍歷二元樹。
- 每個 queue 元素包含 `一個節點` 及 `其對應的剩餘總和`（targetSum - 當前節點值）。
- 如果當前節點是 `葉節點` 且 `剩餘總和為 0`，則 return true。
- 如果 `有子節點`，則將它們 `加入 queue` 並 `更新剩餘總和`。
- 如果遍歷完成後 `沒有找到符合條件的路徑`，則返回 false。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - Method 1. O(n)，其中 n 是二元樹的節點數。
    - Method 2. O(n)，其中 n 是二元樹的節點數。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - Method 1. O(n)，其中 n 是二元樹的節點數，call stack。
    - Method 2. O(n)，其中 n 是二元樹的節點數。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node *TreeNode
	sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Method 1. DFS
	// sum := targetSum - root.Val
	// if sum == 0 && root.Right == nil && root.Left == nil {
	// 	return true
	// }
	// leftTree := hasPathSum(root.Left, sum)
	// rightTree := hasPathSum(root.Right, sum)
	// return leftTree || rightTree

	// Method 2. BFS
	queue := []Node{}
	queue = append(queue, Node{node: root, sum: targetSum - root.Val})

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		node, currSum := curr.node, curr.sum

		if node.Left == nil && node.Right == nil && currSum == 0 {
			return true
		}
		if node.Left != nil {
			queue = append(queue, Node{node: node.Left, sum: currSum - node.Left.Val})
		}
		if node.Right != nil {
			queue = append(queue, Node{node: node.Right, sum: currSum - node.Right.Val})
		}
	}
	return false
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
    private static class Node {
        TreeNode node;
        int sum;

        public Node(TreeNode node, int sum) {
            this.node = node;
            this.sum = sum;
        }
    }

    public boolean hasPathSum(TreeNode root, int targetSum) {
        if (root == null) {
            return false;
        }

        // Method 1. DFS
        // if (root.left == null && root.right == null) {
        //     return root.val == targetSum;
        // }
        // return hasPathSum(root.left, targetSum - root.val) || hasPathSum(root.right, targetSum - root.val);

        // Method 2. BFS
        Queue<Node> queue = new LinkedList<>();
        queue.add(new Node(root, targetSum - root.val));

        while (!queue.isEmpty()) {
            Node current = queue.poll();
            TreeNode node = current.node;
            int currSum = current.sum;

            if (node.left == null && node.right == null && currSum == 0) {
                return true;
            }
            if (node.left != null) {
                queue.add(new Node(node.left, currSum - node.left.val));
            }
            if (node.right != null) {
                queue.add(new Node(node.right, currSum - node.right.val));
            }
        }

        return false;
    }
}
```