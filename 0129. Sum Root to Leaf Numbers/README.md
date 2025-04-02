# Intuition

這題的目標是計算所有 `根到葉子節點的數字總和`，每條路徑上的數字是透過連接節點數值形成的。

例如：
```
    1
   / \
  2   3
```

兩條路徑 `1->2` 和 `1->3` 代表數字 `12` 和 `13`，總和為 `12 + 13 = 25`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

我們可以使用 深度優先搜尋 (DFS) 或 廣度優先搜尋 (BFS) 來解決這個問題。

## Method 1. DFS (遞迴)

- 維護一個變數 sum，用來 `累加當前路徑的數值`。
- 在每個節點，我們將 `sum * 10`，然後 `加上該節點的數值`。
- 當我們 `抵達葉子節點` 時，return sum。
- 最後 return `左右子樹的總和`。

## Method 2. BFS (迴圈)

- 使用 queue 來進行層序遍歷。
- 每當從 queue `取出一個節點` 時，`檢查是否為葉子節點`，若 `是` 則將 `當前數值加到 sum`。
- 若有子節點，則將 `更新後的數值加入 queue`，繼續遍歷。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - Method 1. O(n)，其中 n 是節點數量，每個節點都只會被訪問一次。
    - Method 2. O(n)，其中 n 是節點數量，每個節點都只會被訪問一次。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - Method 1. O(h)，其中 h 是樹的高度，最差情況為 O(n) (call stack)。
    - Method 2. O(n) 空間來存儲 queue 中的節點。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	// return dfs(root, 0)
	return bfs(root)
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	sum := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil && node.Right == nil {
			sum += node.Val
		}
		if node.Left != nil {
			node.Left.Val = add(node.Val, node.Left.Val)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Val = add(node.Val, node.Right.Val)
			queue = append(queue, node.Right)
		}
	}
	return sum
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return add(sum, root.Val)
	}

	leftSum := dfs(root.Left, add(sum, root.Val))
	rightSum := dfs(root.Right, add(sum, root.Val))
	return leftSum + rightSum
}

func add(sum, val int) int {
	return sum*10 + val
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
    public int sumNumbers(TreeNode root) {
        // return dfs(root, 0);
        return bfs(root);
    }

    private int bfs(TreeNode root) {
        if (root == null) return 0;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        int res = 0;
        while (!queue.isEmpty()) {
            TreeNode curNode = queue.poll();
            if (curNode.left == null && curNode.right == null) {
                res += curNode.val;
            }
            if (curNode.left != null) {
                curNode.left.val = add(curNode.val, curNode.left.val);
                queue.add(curNode.left);
            }
            if (curNode.right != null) {
                curNode.right.val = add(curNode.val, curNode.right.val);
                queue.add(curNode.right);
            }
        }
        return res;
    }

    private int dfs(TreeNode root, int sum) {
        if (root == null) return 0;
        if (root.left == null && root.right == null) return add(sum, root.val);
        return dfs(root.left, add(sum, root.val)) + dfs(root.right, add(sum, root.val));
    }

    private int add(int sum, int val) {
        return sum * 10 + val;
    }
}
```