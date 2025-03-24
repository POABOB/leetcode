package leetcode

/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	// Method 1. 全部走過就是樹的節點數
	// if root == nil {
	// 	return 0
	// }

	// left := countNodes(root.Left) + 1
	// right := countNodes(root.Right) + 1

	// return left + right - 1 // root 重複

	// Method 2. 完全二元樹節點數為 2^h - 1
	if root == nil {
		return 0
	}
	tempLeft, tempRight := root, root
	LeftHeight, RightHeight := 0, 0
	for tempLeft != nil {
		tempLeft = tempLeft.Left
		LeftHeight++
	}
	for tempRight != nil {
		tempRight = tempRight.Right
		RightHeight++
	}
	if LeftHeight == RightHeight {
		return (1 << LeftHeight) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)

}

// @lc code=end
