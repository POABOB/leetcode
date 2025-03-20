package leetcode

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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

type Node struct {
	node *TreeNode
	sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Method 1. DFS
	// sum := targetSum - root.Val
	// if sum == 0 && root.Right == nil && root.Left == nil {
	// 	return true
	// }
	// leftTree := hasPathSum(root.Left, sum)
	// rightTree := hasPathSum(root.Right, sum)
	// return leftTree || rightTree

	// Method 2. BFS
	queue := []Node{}
	queue = append(queue, Node{node: root, sum: targetSum - root.Val})

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		node, currSum := curr.node, curr.sum

		if node.Left == nil && node.Right == nil && currSum == 0 {
			return true
		}
		if node.Left != nil {
			queue = append(queue, Node{node: node.Left, sum: currSum - node.Left.Val})
		}
		if node.Right != nil {
			queue = append(queue, Node{node: node.Right, sum: currSum - node.Right.Val})
		}
	}
	return false
}

// @lc code=end
