# Intuition

給定一個 `遞增排序的整數陣列`，將其轉換為 `高度平衡的二元搜尋樹` (Binary Search Tree, BST)。

- 高度平衡的 BST 指的是對於任意節點，`左右子樹的高度`差不超過 `1`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 因為陣列是有被排序過的，可以找出陣列 `中間位置 mid` 作為 root
- 左側子陣列當作 `左子樹`，並`遞迴分割`。
- 右側子陣列當作 `右子樹`，並`遞迴分割`。

```
[1, 2, 3, 4, 5, 6, 7]

Step 1. [1, 2, 3], 4, [5, 6, 7]                 // 4 為 root，left 是 [1, 2, 3]，right 是 [5, 6, 7]

=>                4
              /       \
         [1, 2, 3]  [5, 6, 7]
         
Step 2. left: [1], 2, [3]； right: [5], 6, [7]   // 2 為 root，left 是 [1]，right 是 [3]；6 為 root，left 是 [5]，right 是 [7]

=>                4
              /       \
             2          6
           /   \      /   \
          1     3    5     7
```

<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(n)，每個元素只訪問一次，構建一次節點。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(log n)，遞迴深度為樹高，平衡 BST 高度為 log n。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return process(&nums, 0, len(nums)-1)
}

func process(nums *[]int, start int, end int) *TreeNode {
	mid := start + (end-start)>>1
	root := &TreeNode{Val: (*nums)[mid]}
	root.Left = process(nums, start, mid-1)
	root.Right = process(nums, mid+1, end)
	return root
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
    public TreeNode sortedArrayToBST(int[] nums) {
        return process(nums, 0, nums.length - 1);
    }

    private TreeNode process(int[] nums, int start, int end) {
        if (start > end) {
            return null;
        }
        int mid = start + (end - start) / 2;
        TreeNode root = new TreeNode(nums[mid]);
        root.left = process(nums, start, mid - 1);
        root.right = process(nums, mid + 1, end);
        return root;
    }
}
```