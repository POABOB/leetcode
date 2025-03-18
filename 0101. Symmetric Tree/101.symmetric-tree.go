package leetcode

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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

func isSymmetric(root *TreeNode) bool {
	// Method 1. DFS
	// if root == nil {
	// 	return true
	// }
	// return isMirror(root.Left, root.Right)

	// Method 2. BFS
	queue := make([][2]*TreeNode, 1)
	queue = append(queue, [2]*TreeNode{root.Left, root.Right})
	for len(queue) > 0 {
		leftNode := queue[0][0]
		rightNode := queue[0][1]
		queue = queue[1:]

		if leftNode == nil && rightNode == nil {
			continue
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}

		queue = append(queue, [2]*TreeNode{leftNode.Left, rightNode.Right})
		queue = append(queue, [2]*TreeNode{rightNode.Left, leftNode.Right})
	}
	return true
}

// DFS
// func isMirror(n1 *TreeNode, n2 *TreeNode) bool {
// 	if n1 == nil && n2 == nil {
// 		return true
// 	}
// 	if n1 == nil || n2 == nil {
// 		return false
// 	}
// 	return (n1.Val == n2.Val && isMirror(n1.Left, n2.Right) && isMirror(n1.Right, n2.Left))
// }

// @lc code=end
