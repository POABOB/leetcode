# Intuition

這題的目標是找到二元樹中 `最大路徑總和 (Maximum Path Sum)`。

這個問題較為困難，因為路徑 `不一定要經過根節點`，也不必從葉子節點 `開始` 或 `結束`。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

可以使用 `遞迴 DFS (Depth-First Search)` 來遍歷樹，並在 `過程中記錄最大路徑總和`。

1. 定義遞迴函式 dfs(root)
    - 計算從 `root` 為起點的 `最大單邊路徑和`。
    - 記錄 `全域變數 sum`，用來存儲 `全局最大路徑和`。
2. 計算左子樹與右子樹的最大路徑和
    - 若某條路徑的 `總和為負`數，則 `直接忽略該路徑`，視其值為 `0`。
    - 使用 max(0, dfs(root.Left)) 和 max(0, dfs(root.Right)) 避免負數影響結果。
3. 更新最大路徑和
    - 當前節點 root 可以作為 `「轉折點」`，形成 leftSum + rightSum + root.Val 的完整路徑。
    - 若這條 `完整路徑的總和` 大於 `sum`，則 `更新 sum`。
4. 返回當前節點的最大單邊路徑和
    - 只能選擇 `一條路徑返回給父節點`，不能 `同時選擇左右子樹`，因為二叉樹的路徑必須是單一路徑。
    - 因此，回傳 max(leftSum, rightSum) + root.Val。

## Edge Cases

1. `所有節點都是負數`
    - 由於我們在 max(0, dfs(root.Left)) 過濾掉負數路徑，可能導致 sum 的初始值 `math.MinInt32` 成為最大值。
    - 但由於我們在每個節點都會更新 sum，所以最終結果仍然是單一節點的最大值。
2. 只有 `一個節點` 的情況
    - 如果只有一個節點，則該節點值即為答案。
3. `樹為 nil (空樹)`
    - 這題不會遇到空樹，因為 root 保證存在。

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(n)，因為我們訪問每個節點一次。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(h)，其中 h 是樹的高度，最壞情況下 (不平衡樹) 空間複雜度為 O(n)。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	sum := math.MinInt32
	dfs(root, &sum)
	return sum
}

func dfs(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}

	// 遞迴找尋左右子樹最大 path sum，如果路徑是負，則直接為 0
	leftSum := max(0, dfs(root.Left, sum))
	rightSum := max(0, dfs(root.Right, sum))
	// 每一層都可以進行更新，因為可以不用一定要經過根節點
	*sum = max(leftSum+rightSum+root.Val, *sum)
	// 因為要找最大 path sum，所以只能將兩子樹取最大 + root.Val
	return max(leftSum, rightSum) + root.Val
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
   private int max = Integer.MIN_VALUE;

   public int maxPathSum(TreeNode root) {
      dfs(root);
      return max;
   }

   private int dfs(TreeNode root) {
      if (root == null) return 0;
      int left = Math.max(0, dfs(root.left));
      int right = Math.max(0, dfs(root.right));
      max = Math.max(max, root.val + left + right);
      return root.val + Math.max(left, right);
   }
}
```