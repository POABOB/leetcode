package leetcode

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	leftNode := root.Left
	rightNode := root.Right
	root.Left = nil
	root.Right = leftNode

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = rightNode
}

// @lc code=end
