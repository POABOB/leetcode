# Intuition

這題要求判斷兩棵二元樹是否 `結構相同` 且 `對應節點值相等`。可以使用 `深度優先搜尋 (DFS)` 或 `廣度優先搜尋 (BFS)` 來解這題。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

## 深度優先搜尋 (DFS - 遞迴)

- 如果 p 和 q `同時為 nil`，代表 `兩棵樹都是空的`，`return true`。
- 如果 `其中一棵樹是 nil` 或 `對應節點的值不同`，則 `return false`。
- 遞歸比較 `p.Left 和 q.Left` 以及 `p.Right 和 q.Right`。

## 廣度優先搜尋 (BFS - 層序遍歷)

- 使用 `兩個隊列` 來進行層序遍歷 (BFS)，分別存放 p 和 q 的節點。
- 每次從兩個隊列中取出節點並比較：
    - 如果 `兩個節點都為 nil`，則 `繼續` 遍歷。
    - 如果 `只有一個 nil` 或 `值不同`，則 `return false`。
    - 否則，將 `左右子節點` 分別加入兩個隊列，`繼續比較`。
- 若最終 `沒有發現不匹配` 的情況，則 `return true`。

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Method 1. DFS
	// if p == nil && q == nil {
	// 	return true
	// }
	// if p == nil || q == nil || p.Val != q.Val {
	// 	return false
	// }
	// return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

	// Method 2. BFS
	queue1, queue2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	queue1 = append(queue1, p)
	queue2 = append(queue2, q)
	for len(queue1) > 0 {
		node1 := queue1[0]
		node2 := queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue1 = append(queue1, []*TreeNode{node1.Left, node1.Right}...)
		queue2 = append(queue2, []*TreeNode{node2.Left, node2.Right}...)
	}
	return true
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
    public boolean isSameTree(TreeNode p, TreeNode q) {
        // Method 1. DFS
        // if (p == null && q == null) {
        //     return true;
        // }
        // if (p == null || q == null || p.val != q.val) {
        //     return false;
        // }
        // return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);

        // Method 2. BFS
        Queue<TreeNode> queueP = new LinkedList<>();
        Queue<TreeNode> queueQ = new LinkedList<>();
        queueP.add(p);
        queueQ.add(q);
        while (!queueP.isEmpty()) {
            TreeNode nodeP = queueP.poll();
            TreeNode nodeQ = queueQ.poll();
            if (nodeP == null && nodeQ == null) {
                continue;
            }
            if (nodeP == null || nodeQ == null || nodeP.val != nodeQ.val) {
                return false;
            }
            queueP.add(nodeP.left);
            queueP.add(nodeP.right);
            queueQ.add(nodeQ.left);
            queueQ.add(nodeQ.right);
        }
        return true;
    }
}
```