package leetcode

/*
 * @lc app=leetcode id=427 lang=golang
 *
 * [427] Construct Quad Tree
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
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

// @lc code=end
