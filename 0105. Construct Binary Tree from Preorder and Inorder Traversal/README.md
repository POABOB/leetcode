# Intuition

給定一個二元樹的 `前序遍歷 (Preorder Traversal)` 和 `中序遍歷 (Inorder Traversal)` 的結果，請根據這兩個遍歷結果構建該二元樹。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

1. 前序與中序遍歷特性
    - Preorder Traversal: 根 -> 左子樹 -> 右子樹
    - Inorder Traversal: 左子樹 -> 根 -> 右子樹
- Preorder 的第一個元素是樹的 `根節點`，`
- 根據這個值` 可以在 Inorder 中 `找到該節點的位置`
- 將 Inorder 分成 `左子樹` 和 `右子樹`，然後 `遞迴構建二元樹`。

2. 解法
    - 建立 `Hashmap` 儲存 Inorder，`key` 為 `元素`，`value` 為 `index`，`加速查找`根節點。
    - 使用遞迴構建二元樹，依次處理 `Inorder` 中的元素。
<!-- Describe your approach to solving the problem. -->

# Complexity

- Time complexity
    - O(n): 每個節點遍歷一次。

<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity
    - O(n): Hashmap 需要存儲 O(n) 個節點索引。

<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code

```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	hashmap := make(map[int]int, len(inorder))
	for k, v := range inorder {
		hashmap[v] = k
	}
	return build(&preorder, &inorder, hashmap, 0, len(inorder)-1)
}

func build(preorder *[]int, inorder *[]int, hashmap map[int]int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	// preorder 的 root 統一在首位
	rootValue := (*preorder)[0]
	(*preorder) = (*preorder)[1:]
	// 找出 inorder 內 root 的 index
	i := hashmap[rootValue]
	root := &TreeNode{Val: rootValue}
	root.Left = build(preorder, inorder, hashmap, inStart, i-1)
	root.Right = build(preorder, inorder, hashmap, i+1, inEnd)
	return root
}
```

```java
import java.util.HashMap;
import java.util.Map;

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

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        Map<Integer, Integer> hashmap = new HashMap<>();
        for (int i = 0; i < inorder.length; i++) {
            hashmap.put(inorder[i], i);
        }
        return build(preorder, inorder, hashmap, 0, preorder.length - 1, 0, inorder.length - 1);
    }

    public TreeNode build(int[] preorder, int[] inorder, Map<Integer, Integer> hashmap, int preStart, int preEnd, int inStart, int inEnd) {
        if (preStart > preEnd) {
            return null;
        }
        int rootValue = preorder[preStart];
        int i = hashmap.get(rootValue);
        TreeNode root = new TreeNode(rootValue);
        root.left = build(preorder, inorder, hashmap, preStart + 1, preStart + i - inStart, inStart, i - 1);
        root.right = build(preorder, inorder, hashmap, preStart + i - inStart + 1, preEnd, i + 1, inEnd);
        return root;
    }
}
```