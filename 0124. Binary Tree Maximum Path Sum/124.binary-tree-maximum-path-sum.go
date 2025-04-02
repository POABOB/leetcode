package leetcode

import "math"

/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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

// @lc code=end
