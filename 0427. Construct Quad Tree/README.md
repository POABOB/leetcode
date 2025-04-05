# Intuition
給定一個 `n x n` 的二維二元矩陣 grid，每個值為 `0` 或 `1`，請構造一棵 `對應的四元樹 (Quad Tree)`。

- 四元樹是一種節點最多有四個子節點的樹，常用於 `區域劃分` 和 `圖像壓縮`。
- 構建的原則是：
	- 若一塊區域中 `所有值相同 (全 0 或全 1)`，則為 `葉節點 (Leaf)`，並設定 `IsLeaf = true` 與對應的值。
	- 否則繼續將該區域等分為四個子區域，`遞迴構建四個子節點`。

<!-- Describe your first thoughts on how to solve this problem. -->

# Approach

- 遞迴劃分區域：
	- 對於一個區域 (x, y, width)，若 width == 1，表示 `只剩單一格`，直接 `回傳對應的 Node`。
	- 否則將該區域 `切分` 為：
		- TopLeft：`左上`
		- TopRight：`右上`
		- BottomLeft：`左下`
		- BottomRight：`右下`
	- 對四個子區域 `遞迴進行相同處理`。
- 合併邏輯 (壓縮)：
	- 若四個子區域 `皆為葉節點`，且 `值相同`，則 `可壓縮為一個葉節點`。
- 傳參：
	- 使用 `grid *[][]int` 傳入指標 `避免 slice 複製開銷`。
	- 利用 (x, y) 表示 `左上角座標`，w 表示當前 `區域的寬度`。
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity
    - O(n＾2)：每個格子最多被訪問一次，節點數最多為 O(n＾2)。
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity 
    - O(n＾2)：最壞情況下，無法合併節點，導致建構出最多的樹節點。
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```go
package leetcode

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return dfs(0, 0, len(grid[0]), &grid)
}

func dfs(x, y, w int, grid *[][]int) *Node {
	if w == 1 {
		return &Node{Val: (*grid)[x][y] == 1, IsLeaf: true}
	}

	// Divide and Conquer
	halfW := w >> 1
	topLeft := dfs(x, y, halfW, grid)
	topRight := dfs(x+halfW, y, halfW, grid)
	bottomLeft := dfs(x, y+halfW, halfW, grid)
	bottomRight := dfs(x+halfW, y+halfW, halfW, grid)

	// chceck leaf and value
	node := &Node{}
	// chceck leaf and value
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf && topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		node.Val = topLeft.Val
		node.IsLeaf = true
	} else {
		node.TopLeft = topLeft
		node.TopRight = topRight
		node.BottomLeft = bottomLeft
		node.BottomRight = bottomRight
	}
	return node
}
```

```java
class Node {
    public boolean val;
    public boolean isLeaf;
    public Node topLeft;
    public Node topRight;
    public Node bottomLeft;
    public Node bottomRight;


    public Node() {
        this.val = false;
        this.isLeaf = false;
        this.topLeft = null;
        this.topRight = null;
        this.bottomLeft = null;
        this.bottomRight = null;
    }

    public Node(boolean val, boolean isLeaf) {
        this.val = val;
        this.isLeaf = isLeaf;
        this.topLeft = null;
        this.topRight = null;
        this.bottomLeft = null;
        this.bottomRight = null;
    }

    public Node(boolean val, boolean isLeaf, Node topLeft, Node topRight, Node bottomLeft, Node bottomRight) {
        this.val = val;
        this.isLeaf = isLeaf;
        this.topLeft = topLeft;
        this.topRight = topRight;
        this.bottomLeft = bottomLeft;
        this.bottomRight = bottomRight;
    }
}

class Solution {
    public Node construct(int[][] grid) {
        return dfs(0, 0, grid.length, grid);
    }

    private Node dfs(int x, int y, int w, int[][] grid) {
        if (w == 1) return new Node(grid[y][x] == 1, true);

        int halfW = w >> 1;
        Node topLeft = dfs(x, y, halfW, grid);
        Node topRight = dfs(x + halfW, y, halfW, grid);
        Node bottomLeft = dfs(x, y + halfW, halfW, grid);
        Node bottomRight = dfs(x + halfW, y + halfW, halfW, grid);

        Node node = new Node();
        if (topRight.isLeaf && topLeft.isLeaf && bottomLeft.isLeaf && bottomRight.isLeaf && topLeft.val == topRight.val && topLeft.val == bottomLeft.val && topLeft.val == bottomRight.val) {
            node.val = topLeft.val;
            node.isLeaf = true;
        } else {
            node.topLeft = topLeft;
            node.topRight = topRight;
            node.bottomLeft = bottomLeft;
            node.bottomRight = bottomRight;
        }
        return node;
    }
}
```