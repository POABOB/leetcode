# Intuition
這題的目標是設計一個 `二元搜尋樹 (BST)` 迭代器，可以按照遞增順序遍歷 BST，並且確保 Next() 和 HasNext() 操作的效率足夠高。
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 我們使用 `中序遍歷 (Inorder Traversal)` 來將 BST 轉換成遞增排序的陣列。
- Constructor(root) 會執行一次中序遍歷，並 `將所有節點存入 inorder 陣列`。
- Next() 每次 `返回當前索引的節點值`，然後 `將索引向後移動`。
- HasNext() 判斷 `當前索引是否仍在數組範圍內`。
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - `Constructor()` 需要 O(n) 時間來遍歷整棵樹，其中 n 是 BST 的節點數量。
    - `Next()` 和 `HasNext()` 都是 O(1) 時間複雜度，因為它們只是讀取陣列。
- Space complexity 
    - O(n)，因為我們用了一個陣列來存儲 BST 所有節點的中序遍歷結果。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	inorder   []*TreeNode
	currIndex int
}

func Constructor(root *TreeNode) BSTIterator {
	arr := make([]*TreeNode, 0)
	inorder(root, &arr)
	return BSTIterator{inorder: arr, currIndex: 0}
}

func inorder(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root)
	inorder(root.Right, res)
}

func (this *BSTIterator) Next() int {
	curr := this.inorder[this.currIndex]
	this.currIndex++
	return curr.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.currIndex < len(this.inorder)
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

class BSTIterator {

    private final List<TreeNode> inorderList = new ArrayList<>();
    private int index = 0;

    public BSTIterator(TreeNode root) {
        inorder(root);
    }

    private void inorder(TreeNode root) {
        if (root == null) return;
        inorder(root.left);
        inorderList.add(root);
        inorder(root.right);
    }

    public int next() {
        return inorderList.get(index++).val;
    }

    public boolean hasNext() {
        return index < inorderList.size();
    }
}
```
