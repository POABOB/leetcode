# Intuition

給定一個二元樹的 `中序遍歷 (Inorder Traversal)` 和 `後序遍歷 (Postorder Traversal)` 的結果，請根據這兩個遍歷結果構建該二元樹。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

1. 中序與後序遍歷特性
   - Inorder Traversal: 左子樹 -> 根 -> 右子樹
   - Postorder Traversal: 左子樹 -> 右子樹 -> 根
   - Postorder 的 `最後一個元素` 是樹的 `根節點`，根據這個值可以在 `Inorder` 中找到該節點的位置
   - 將 `Inorder` 分成左子樹部分和右子樹部分，然後遞迴構建二元樹。
2. 解法
   - 建立 Hashmap 儲存 Inorder，`key` 為 `元素`，`value` 為 `index`，`加速查找`根節點索引。
   - 使用遞迴構造二元樹，依次處理 Postorder 中的元素。
   - 由於後序遍歷的 `最後一個元素` 是 `根節點`，因此我們應當先構造 `右子樹`，再構造 `左子樹`。
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 建立 hashmap 儲存 inorder value to key
	hashmap := make(map[int]int, len(inorder))
	for k, v := range inorder {
		hashmap[v] = k
	}
	return build(&inorder, &postorder, hashmap, 0, len(inorder)-1)
}

func build(inorder *[]int, postorder *[]int, hashmap map[int]int, inStart, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	// postorder 要從元素最後一位取
	postLen := len((*postorder))
	rootValue := (*postorder)[postLen-1]
	(*postorder) = (*postorder)[:postLen-1]
	index := hashmap[rootValue]
	root := &TreeNode{Val: rootValue}
	// 這邊會先建構右子樹，原因是 (*postorder)[postLen-1] 每次都取最後一位
	// postorder 內 [{左},{右},root] 的結構
	// 如果由後面往前取的話，那麼每次找到的跟都會先是右子樹的根
	root.Right = build(inorder, postorder, hashmap, index+1, inEnd)
	root.Left = build(inorder, postorder, hashmap, inStart, index-1)
	return root
}
```

```java
import java.util.HashMap;

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
    // 存储 inorder 中值到索引的映射
    HashMap<Integer, Integer> valToIndex = new HashMap<>();

    public TreeNode buildTree(int[] inorder, int[] postorder) {
        for (int i = 0; i < inorder.length; i++) {
            valToIndex.put(inorder[i], i);
        }
        return build(inorder, 0, inorder.length - 1, postorder, 0, postorder.length - 1);
    }

    private TreeNode build(int[] inorder, int inStart, int inEnd, int[] postorder, int postStart, int postEnd) {
        if (inStart > inEnd) {
            return null;
        }
        int rootVal = postorder[postEnd];
        int index = valToIndex.get(rootVal);
        int leftSize = index - inStart;
        TreeNode root = new TreeNode(rootVal);

        root.left = build(inorder, inStart, index - 1,
                postorder, postStart, postStart + leftSize - 1);

        root.right = build(inorder, index + 1, inEnd,
                postorder, postStart + leftSize, postEnd - 1);
        return root;
    }
}
```